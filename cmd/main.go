package main

import (
	"SuperStub/internal/config"
	"SuperStub/internal/dynamic"
	g "SuperStub/internal/services/grpc"
	"SuperStub/internal/services/project"
	r "SuperStub/internal/services/rest"
	"SuperStub/internal/storage/postgresql"
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jhump/protoreflect/desc/protoparse"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/dynamicpb"
	"io"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres_user"
	password = "postgres_password"
	dbname   = "postgres_db"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	logger := setupLogger(cfg.Env)

	storage := setupStorage()

	srvRest := startRestServer(storage, logger)
	srvGrpc := startGrpcServer(storage, logger, cfg.GRPC)

	gracefulShutdown(srvRest, srvGrpc, logger)
}

func gracefulShutdown(srv *http.Server, srvGrpc *grpc.Server, log *slog.Logger) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	err := srv.Shutdown(context.TODO())
	if err != nil {
		log.Error(err.Error())
		return
	}

	srvGrpc.GracefulStop()

	log.Info("Gracefully stopped")
}

func startGrpcServer(storage *postgresql.Storage, log *slog.Logger, cfg config.GRPCConfig) *grpc.Server {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Error("failed to listen: %v", err)
		panic(err)
	}

	newService := dynamic.NewService("dynamic.Service")

	interceptor := grpc.UnknownServiceHandler(func(srv any, stream grpc.ServerStream) error {
		log.Debug("UnknownServiceHandler called")

		sm, ok := grpc.MethodFromServerStream(stream)
		if !ok {
			return errors.New("failed to get stream method")
		}
		if sm != "" && sm[0] == '/' {
			sm = sm[1:]
		}
		pos := strings.LastIndex(sm, "/")
		if pos == -1 {
			errMsg := fmt.Sprintf("malformed method name %q", sm)
			return errors.New(errMsg)
		}
		service := sm[:pos]
		method := sm[pos+1:]
		protoName := service[:strings.LastIndex(sm, ".")]
		log.Info(fmt.Sprintf("service: %s, method: %s, proto: %s", service, method, protoName))

		protoString, err := storage.GetProtoByName(context.TODO(), protoName)
		if err != nil {
			log.Error("Failed to parse proto files: %v", err)
			return err
		}

		fileName := fmt.Sprintf("%s.proto", protoName)
		err = createProtoFile(fileName, protoString)
		if err != nil {
			log.Error("Failed to parse proto files: %v", err)
			return err
		}
		defer os.Remove(fileName)

		protoFiles := []string{fmt.Sprintf("%s.proto", protoName)}
		parser := protoparse.Parser{}
		fileDescriptors, err := parser.ParseFiles(protoFiles...)
		if err != nil {
			log.Error("Failed to parse proto files: %v", err)
			return err
		}

		serviceFd := fileDescriptors[0].FindService(service)
		methodFd := serviceFd.FindMethodByName(method)
		rqFdName := methodFd.GetInputType()
		rsFdName := methodFd.GetOutputType()
		messageRq := dynamicpb.NewMessage(rqFdName.UnwrapMessage())
		messageRs := dynamicpb.NewMessage(rsFdName.UnwrapMessage())

		if methodFd.IsClientStreaming() || methodFd.IsServerStreaming() {
			//TODO separate unary and streaming
		}

		if err := stream.RecvMsg(messageRq); err != nil {
			return err
		}

		stub, err := storage.GetGrpcStubByFullPath(context.TODO(), fmt.Sprintf("%s.%s", service, method))
		if err != nil {
			log.Error("Failed to get grpc stub: %v", err)
			return err
		}
		jsonData := []byte(stub.ResponseBody)

		unmarshaler := protojson.UnmarshalOptions{
			DiscardUnknown: true,
		}
		if err := unmarshaler.Unmarshal(jsonData, messageRs); err != nil {
			log.Error("Failed to unmarshal JSON: %v", err)
			return err
		}

		err = stream.SendMsg(messageRs)
		if err != nil {
			return err
		}

		return nil
	})

	s := dynamic.NewServer([]*dynamic.Service{newService}, interceptor)
	log.Info("Starting grpc server listening at " + lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		log.Error("failed to serve: %v", err)
		panic(err)
	}
	return s
}

func createProtoFile(filename, protoString string) error {
	dst, err := os.Create(filename)
	if err != nil {
		log.Println("error creating file", err)
		return err
	}
	defer dst.Close()
	if _, err := io.Copy(dst, strings.NewReader(protoString)); err != nil {
		return err
	}
	return nil
}

func startRestServer(storage *postgresql.Storage, log *slog.Logger) *http.Server {
	router := setupRouter(storage)

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		log.Info("Starting server on http://" + srv.Addr)
		err := srv.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()
	return srv
}

func setupRouter(storage *postgresql.Storage) *mux.Router {
	projectService := project.New(nil, storage, storage, storage, storage)
	restService := r.New(nil, storage, storage, storage, storage)
	grpcService := g.New(nil, storage, storage, storage, storage, storage)

	router := mux.NewRouter()

	router.HandleFunc("/projects", projectService.GetAll).Methods("GET")
	router.HandleFunc("/projects/{project_id}", projectService.GetById).Methods("GET")
	router.HandleFunc("/projects", projectService.Create).Methods("POST")

	router.HandleFunc("/projects/{project_id}/stub", restService.GetAllRestStubs).Methods("GET")
	router.HandleFunc("/projects/{project_id}/stub/{id}", restService.GetRestStubById).Methods("GET")
	router.HandleFunc("/projects/{project_id}/stub", restService.CreateRestStub).Methods("POST")
	router.HandleFunc("/projects/{project_id}/stub/{id}", restService.UpdateRestStub).Methods("PUT")
	router.HandleFunc("/projects/{project_id}/stub/{id}", restService.DeleteRestStub).Methods("DELETE")

	router.HandleFunc("/projects/{project_id}/{path}", restService.ServeStub).Methods("GET")

	router.HandleFunc("/projects/{project_id}/grpc/upload-proto", grpcService.UploadProto).Methods("POST")

	router.HandleFunc("/projects/{project_id}/grpc/stub", grpcService.GetAllGrpcStubs).Methods("GET")
	router.HandleFunc("/projects/{project_id}/grpc/stub/{id}", grpcService.GetGrpcStubById).Methods("GET")
	router.HandleFunc("/projects/{project_id}/grpc/stub", grpcService.CreateGrpcStub).Methods("POST")
	router.HandleFunc("/projects/{project_id}/grpc/stub/{id}", grpcService.UpdateGrpcStub).Methods("PUT")
	router.HandleFunc("/projects/{project_id}/grpc/stub/{id}", grpcService.DeleteGrpcStub).Methods("DELETE")
	return router
}

func setupStorage() *postgresql.Storage {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	storage, err := postgresql.New(psqlInfo)
	if err != nil {
		panic(err)
	}
	return storage
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
