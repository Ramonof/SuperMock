package main

//func main() {
//
//	protoFiles := []string{"example.proto"}
//
//	parser := protoparse.Parser{}
//
//	fileDescriptors, err := parser.ParseFiles(protoFiles...)
//	if err != nil {
//		log.Fatalf("Failed to parse proto files: %v", err)
//	}
//
//	for _, fd := range fileDescriptors {
//		fdProto := fd.AsFileDescriptorProto()
//
//		fmt.Printf("Proto file name: %s\n", fdProto.GetName())
//
//		for _, message := range fdProto.GetMessageType() {
//			fmt.Printf("Message type: %s\n", message.GetName())
//		}
//	}
//}

//func main() {
//	// Read the generated descriptor.pb file
//	descriptorFile := "descriptor.pb"
//	descriptorContent, err := ioutil.ReadFile(descriptorFile)
//	if err != nil {
//		log.Fatalf("Failed to read descriptor file: %v", err)
//	}
//
//	// Check if the file is not empty
//	if len(descriptorContent) == 0 {
//		log.Fatalf("Descriptor file is empty")
//	}
//
//	// Unmarshal the descriptor into a FileDescriptorSet
//	var fileDescriptorSet descriptorpb.FileDescriptorSet
//	if err := proto.Unmarshal(descriptorContent, &fileDescriptorSet); err != nil {
//		log.Fatalf("Failed to unmarshal descriptor file: %v", err)
//	}
//
//	// Access the first FileDescriptorProto
//	if len(fileDescriptorSet.File) > 0 {
//		fileDescriptorProto := fileDescriptorSet.File[0]
//
//		// Print the descriptor proto
//		fmt.Printf("FileDescriptorProto: %v\n", fileDescriptorProto)
//		fmt.Println()
//		fmt.Println(fileDescriptorProto.Service)
//	} else {
//		log.Println("No file descriptors found in the set.")
//	}
//}
