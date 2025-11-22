package handler

import (
	"context"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"

	"github.com/bufbuild/protocompile"
	"github.com/bufbuild/protocompile/protoutil"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/pluginpb"
)

func TestProtoFile(t *testing.T) {
	t.Parallel()

	cwd, err := os.Getwd()
	require.NoError(t, err)

	todoModelProto := "todo/v1/todo_model.proto"
	compiler := protocompile.Compiler{
		Resolver: &protocompile.SourceResolver{
			ImportPaths: []string{
				filepath.Join(cwd, "../../test-protos"),
				filepath.Join(cwd, "../../protos"),
			},
		},
	}

	fileDescriptors, err := compiler.Compile(context.Background(), todoModelProto)
	require.NoError(t, err)

	response := &pluginpb.CodeGeneratorResponse{}

	for _, fileDescriptor := range fileDescriptors {
		descriptor := protoutil.ProtoFromFileDescriptor(fileDescriptor)
		err = ProtoFile(descriptor, response)
		require.NoError(t, err)
	}

	require.Len(t, response.File, 1, "Expected exactly one generated file")

	generatedFile := response.File[0]
	require.Equal(t, "sqlc.yaml", *generatedFile.Name, "Expected generated file to be sqlc.yaml")
	require.NotNil(t, generatedFile.Content, "Generated file content should not be nil")
	assert.Contains(t, *generatedFile.Content, "version: \"2\"", "Expected version 2 in sqlc.yaml")
	assert.Contains(t, *generatedFile.Content, "name: todo", "Expected todo domain in sqlc.yaml")
	assert.Contains(t, *generatedFile.Content, "package: todo_gendb", "Expected todo_gendb package in sqlc.yaml")
}
