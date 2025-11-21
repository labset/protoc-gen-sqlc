package handler

import (
	"context"
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
}
