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

		if methodFd.IsClientStreaming() || methodFd.IsServerStreaming() {
			//TODO
		}

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

//Get the Field Descriptor for the New Field:
//Use messageDescriptor.Fields().ByName("fieldName") or messageDescriptor.Fields().ByNumber(fieldNumber) to get the protoreflect.FieldDescriptor for the specific field you want to set.
//Set the Field Value:
//For scalar fields (e.g., string, int32, bool):
//Go
//
//fieldDesc := messageDescriptor.Fields().ByName("your_scalar_field_name")
//msg.Set(fieldDesc, protoreflect.ValueOf("your_value"))
//For composite fields (e.g., nested messages, repeated fields, maps):
//Nested Message:
//Go
//
//nestedFieldDesc := messageDescriptor.Fields().ByName("your_nested_message_field_name")
//nestedMsg := dynamicpb.NewMessage(nestedFieldDesc.Message()) // Create a new nested message
//// Populate nestedMsg fields
//msg.Set(nestedFieldDesc, protoreflect.ValueOfMessage(nestedMsg.ProtoReflect()))
//repeated fields.
//Go
//
//repeatedFieldDesc := messageDescriptor.Fields().ByName("your_repeated_field_name")
//list := msg.Mutable(repeatedFieldDesc).List()
//list.Append(protoreflect.ValueOf("item1"))
//list.Append(protoreflect.ValueOf("item2"))
//map fields.
//Go
//
//mapFieldDesc := messageDescriptor.Fields().ByName("your_map_field_name")
//m := msg.Mutable(mapFieldDesc).Map()
//m.Set(protoreflect.ValueOf("key1").MapKey(), protoreflect.ValueOf("value1"))

//j := `{	"@type": "echo.EchoRequest", "firstName": "sal", "lastName": "mander", "middleName": { "name": "a"}}`
//a, err := anypb.New(echoRequestMessageType.New().Interface())
//
//err = protojson.Unmarshal([]byte(j), a)
//fmt.Printf("Encoded EchoRequest using protojson and anypb %v\n", hex.EncodeToString(a.Value))

//https://blog.salrashid.dev/articles/2022/grpc_wireformat/
