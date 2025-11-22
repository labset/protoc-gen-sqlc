package handler

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/descriptorpb"

	"github.com/bufbuild/protocompile"
	"github.com/bufbuild/protocompile/protoutil"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/pluginpb"
)

func TestProcessProtoFiles(t *testing.T) {
	t.Parallel()

	cwd, err := os.Getwd()
	require.NoError(t, err)

	compiler := protocompile.Compiler{
		Resolver: &protocompile.SourceResolver{
			ImportPaths: []string{
				filepath.Join(cwd, "../../test-protos"),
				filepath.Join(cwd, "../../protos"),
			},
		},
	}

	todoModelProto := "todo/v1/todo_model.proto"
	musicModelProto := "music/v1/music_model.proto"
	fileDescriptors, err := compiler.Compile(context.Background(), todoModelProto, musicModelProto)
	require.NoError(t, err)

	protoDescriptors := make([]*descriptorpb.FileDescriptorProto, len(fileDescriptors))
	for i, fd := range fileDescriptors {
		protoDescriptors[i] = protoutil.ProtoFromFileDescriptor(fd)
	}

	response := &pluginpb.CodeGeneratorResponse{}

	err = ProcessProtoFiles(protoDescriptors, response)
	require.NoError(t, err)

	require.Len(t, response.GetFile(), 1, "Expected exactly one generated file")

	generatedFile := response.GetFile()[0]
	require.Equal(
		t,
		"sqlc.yaml",
		generatedFile.GetName(),
		"Expected generated file to be sqlc.yaml",
	)
	require.NotNil(t, generatedFile.GetContent(), "Generated file content should not be nil")
	assert.Contains(
		t,
		generatedFile.GetContent(),
		"version: \"2\"",
		"Expected version 2 in sqlc.yaml",
	)
	assert.Contains(
		t,
		generatedFile.GetContent(),
		"name: todo",
		"Expected todo domain in sqlc.yaml",
	)
	assert.Contains(
		t,
		generatedFile.GetContent(),
		"package: gendb_todo",
		"Expected gendb_todo package in sqlc.yaml",
	)
}
