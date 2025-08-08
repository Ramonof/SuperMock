package main

import (
	"SuperStub/cmd/grpc/dynamic"
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
	"log"
	"net"
	"strings"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type serverKey struct{}

func contextWithServer(ctx context.Context, server *grpc.Server) context.Context {
	return context.WithValue(ctx, serverKey{}, server)
}

func printFd(fileDescriptors []*desc.FileDescriptor) {
	for _, fd := range fileDescriptors {
		fdProto := fd.AsFileDescriptorProto()
		fmt.Printf("Proto file name: %s\n", fdProto.GetName())

		for _, message := range fdProto.GetMessageType() {
			fmt.Printf("Message type: %s\n", message.GetName())
		}
	}
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	newService := dynamic.NewService("dynamic.Service")

	interceptor := grpc.UnknownServiceHandler(func(srv any, stream grpc.ServerStream) error {
		//ctx := stream.Context()
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
		log.Printf("service: %s, method: %s, proto: %s\n", service, method, protoName)

		protoFiles := []string{fmt.Sprintf("%s.proto", protoName)}
		parser := protoparse.Parser{}
		fileDescriptors, err := parser.ParseFiles(protoFiles...)
		if err != nil {
			log.Fatalf("Failed to parse proto files: %v", err)
		}

		serviceFd := fileDescriptors[0].FindService(service)
		methodFd := serviceFd.FindMethodByName(method)
		rqFdName := methodFd.GetInputType()
		rsFdName := methodFd.GetOutputType()
		messageRq := dynamicpb.NewMessage(rqFdName.UnwrapMessage())
		messageRs := dynamicpb.NewMessage(rsFdName.UnwrapMessage())

		if err := stream.RecvMsg(messageRq); err != nil {
			return err
		}
		name := messageRq.Get(rqFdName.FindFieldByName("name").UnwrapField())
		rsMessage := fmt.Sprintf("Hello %s!", name)
		messageRs.Set(rsFdName.FindFieldByName("message").UnwrapField(), protoreflect.ValueOfString(rsMessage))

		err = stream.SendMsg(messageRs)
		if err != nil {
			return err
		}

		return nil
	})
	s := dynamic.NewServer([]*dynamic.Service{newService}, interceptor)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
