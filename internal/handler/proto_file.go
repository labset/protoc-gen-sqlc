package handler

import (
	"github.com/labset/protoc-gen-sqlc/internal/codegen"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

func ProtoFile(
	file *descriptorpb.FileDescriptorProto,
	response *pluginpb.CodeGeneratorResponse,
) error {
	messages := file.GetMessageType()
	if len(messages) == 0 {
		return nil
	}

	// handle configuration
	if configFile := codegen.GenerateSqlcConfigFile(file); configFile != nil {
		response.File = append(response.File, configFile)
	}

	// handle migrations (schema)
	// handle queries

	// output structure:
	// sqlc.yaml
	// data/
	//   queries/
	//   migrations/
	return nil
}
