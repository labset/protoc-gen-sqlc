package main

import (
	"github.com/viqueen/go-protoc-gen-plugin/internal/handler"
	"io"
	"log"
	"os"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("failed to read from stdin: %v", err)
	}

	request := &pluginpb.CodeGeneratorRequest{}
	if err = proto.Unmarshal(data, request); err != nil {
		log.Fatalf("failed to unmarshal input: %v", err)
	}

	response := &pluginpb.CodeGeneratorResponse{}
	for _, protoFile := range request.GetProtoFile() {
		if err = handler.ProtoFile(protoFile, response); err != nil {
			log.Fatalf("failed to process proto file: %v", err)
		}
	}

	respond(response)
}

func respond(response *pluginpb.CodeGeneratorResponse) {
	out, err := proto.Marshal(response)
	if err != nil {
		log.Fatalf("Failed to marshal response: %v", err)
	}

	_, err = os.Stdout.Write(out)
	if err != nil {
		log.Fatalf("Failed to write response: %v", err)
	}
}
