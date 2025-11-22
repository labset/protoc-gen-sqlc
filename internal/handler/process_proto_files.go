package handler

import (
	"github.com/labset/protoc-gen-sqlc/internal/codegen"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

func ProcessProtoFiles(
	files []*descriptorpb.FileDescriptorProto,
	response *pluginpb.CodeGeneratorResponse,
) error {
	var domains []string

	for _, file := range files {
		if len(file.GetMessageType()) == 0 {
			continue
		}

		domain := codegen.ExtractDomain(file.GetPackage())
		if domain != "" {
			domains = append(domains, domain)
		}
	}

	if len(domains) > 0 {
		if configFile := codegen.GenerateSqlcConfigFileForDomains(domains); configFile != nil {
			response.File = append(response.File, configFile)
		}
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
