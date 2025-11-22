package handler

import (
	"github.com/labset/protoc-gen-sqlc/internal/codegen"
	"github.com/labset/protoc-gen-sqlc/internal/helpers"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

func ProcessProtoFiles(
	files []*descriptorpb.FileDescriptorProto,
	response *pluginpb.CodeGeneratorResponse,
) error {
	handleConfig(files, response)

	// handle migrations (schema)
	// handle queries

	// output structure:
	// sqlc.yaml
	// data/
	//   queries/
	//   migrations/
	return nil
}

func handleConfig(files []*descriptorpb.FileDescriptorProto, response *pluginpb.CodeGeneratorResponse) {
	domains := make(map[string]bool)

	for _, file := range files {
		if len(file.GetMessageType()) == 0 {
			continue
		}

		domain := helpers.ExtractDomain(file.GetPackage())
		if domain != "" {
			domains[domain] = true
		}
	}

	if len(domains) == 0 {
		return
	}

	if configFile := codegen.SqlcConfigGen(domains); configFile != nil {
		response.File = append(response.File, configFile)
	}
}
