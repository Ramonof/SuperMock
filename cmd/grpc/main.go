package main

import (
	"SuperStub/cmd/grpc/dynamic"
	pb "SuperStub/cmd/grpc/helloworld"
	"context"
	"flag"
	"fmt"
	"github.com/jhump/protoreflect/desc/protoparse"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(_ context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

//func main() {
//	flag.Parse()
//	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
//	//con, _ := lis.Accept()
//	//st := s.newHTTP2Transport(rawConn)
//	//Res, _ := io.ReadAll(con)
//	//fmt.Println(Res)
//	//
//
//	//s.RegisterService()
//	//protoregistry.GlobalFiles.FindFileByPath().RegisterFile()
//
//	if err != nil {
//		log.Fatalf("failed to listen: %v", err)
//	}
//	s := grpc.NewServer()
//	pb.RegisterGreeterServer(s, &server{})
//	log.Printf("server listening at %v", lis.Addr())
//	if err := s.Serve(lis); err != nil {
//		log.Fatalf("failed to serve: %v", err)
//	}
//}

type Req struct {
	Name string
}

type Res struct {
	Message string
}

//func main() {
//	//encoding.RegisterCodec(&dynamic.Codec{})
//	flag.Parse()
//	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
//
//	if err != nil {
//		log.Fatalf("failed to listen: %v", err)
//	}
//
//	service := dynamic.NewService("helloworld.Greeter")
//	service.RegisterUnaryMethod("SayHello", new(dynamicpb.Message), new(Res), func(ctx context.Context, in interface{}) (interface{}, error) {
//		req := in.(*dynamicpb.Message)
//		return &Res{Message: fmt.Sprintf("hi, %s", req)}, nil
//	})
//	s := dynamic.NewServer([]*dynamic.Service{service})
//
//	log.Printf("server listening at %v", lis.Addr())
//	if err := s.Serve(lis); err != nil {
//		log.Fatalf("failed to serve: %v", err)
//	}
//}

//func main() {
//	flag.Parse()
//	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
//	if err != nil {
//		log.Fatalf("failed to listen: %v", err)
//	}
//
//	service := dynamic.NewService("helloworld.Greeter")
//	service.RegisterUnaryMethod("SayHello", new(structpb.Struct), new(pb.HelloReply), func(ctx context.Context, in interface{}) (interface{}, error) {
//		req := in.(*structpb.Struct).Fields
//		res := &structpb.Struct{Fields: map[string]*structpb.Value{}}
//		res.Fields["message"] = structpb.NewStringValue(fmt.Sprintf("hi, %s", req["fields"]))
//		return res, nil
//		//return &pb.HelloReply{Message: "Hello"}, nil
//	})
//	s := dynamic.NewServer([]*dynamic.Service{service})
//
//	log.Printf("server listening at %v", lis.Addr())
//	if err := s.Serve(lis); err != nil {
//		log.Fatalf("failed to serve: %v", err)
//	}
//}

type serverKey struct{}

func contextWithServer(ctx context.Context, server *grpc.Server) context.Context {
	return context.WithValue(ctx, serverKey{}, server)
}

func main() {
	flag.Parse()

	protoFiles := []string{"helloworld.proto"}

	parser := protoparse.Parser{}

	fileDescriptors, err := parser.ParseFiles(protoFiles...)
	if err != nil {
		log.Fatalf("Failed to parse proto files: %v", err)
	}

	for _, fd := range fileDescriptors {
		fdProto := fd.AsFileDescriptorProto()
		fmt.Printf("Proto file name: %s\n", fdProto.GetName())

		for _, message := range fdProto.GetMessageType() {
			fmt.Printf("Message type: %s\n", message.GetName())
		}
	}
	mdRq := fileDescriptors[0].FindMessage("helloworld.HelloRequest")
	messageRq := dynamicpb.NewMessage(mdRq.UnwrapMessage())
	mdRs := fileDescriptors[0].FindMessage("helloworld.HelloReply")
	messageRs := dynamicpb.NewMessage(mdRs.UnwrapMessage())

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	service := dynamic.NewService("helloworld.Greeter")
	service.RegisterUnaryMethod("SayHello2", messageRq, messageRs, func(ctx context.Context, in interface{}) (interface{}, error) {
		out := messageRs.New()
		inTyped := in.(protoreflect.Message)
		//TODO dynamic mdRq?
		name := inTyped.Get(mdRq.FindFieldByName("name").UnwrapField())
		rsMessage := fmt.Sprintf("Hello %s!", name)
		//md.FindFieldByName()
		//out.Set(md.FindFieldByNumber(1).UnwrapField(), protoreflect.ValueOfString("str"))
		out.Set(mdRs.FindFieldByName("message").UnwrapField(), protoreflect.ValueOfString(rsMessage))
		return out, nil
	})
	interceptor := grpc.UnknownServiceHandler(func(srv any, stream grpc.ServerStream) error {
		fmt.Println("Unknown")
		//ctx := stream.Context()

		////m := new(dynamicpb.Message)
		//parse(stream, messageRq)
		//im := proto.Message(m)
		if err := stream.RecvMsg(messageRq); err != nil {
			return err
		}
		name := messageRq.Get(mdRq.FindFieldByName("name").UnwrapField())
		rsMessage := fmt.Sprintf("Hello %s!", name)
		messageRs.Set(mdRs.FindFieldByName("message").UnwrapField(), protoreflect.ValueOfString(rsMessage))

		err := stream.SendMsg(messageRs)
		if err != nil {
			return err
		}

		return nil
	})
	s := dynamic.NewServer([]*dynamic.Service{service}, interceptor)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func parse(stream grpc.ServerStream, m proto.Message) {
	if err := stream.RecvMsg(m); err != nil {
		panic(err)
	}
}
