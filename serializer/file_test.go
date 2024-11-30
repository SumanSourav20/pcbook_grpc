package serializer_test

import (
	"testing"

	"github.com/SumanSourav20/pcbook_grpc/pb"
	"github.com/SumanSourav20/pcbook_grpc/sample"
	"github.com/SumanSourav20/pcbook_grpc/serializer"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestWriteProtobufToBinaryFile(t *testing.T) {

	laptop := sample.NewLaptop()

	err := serializer.WriteProtobufToBinaryFile(laptop, "new_laptop.bin")
	require.NoError(t, err)

	var newLaptop pb.Laptop

	err = serializer.ReadProtobufFromBinaryFile("new_laptop.bin", &newLaptop)
	require.NoError(t, err)

	require.True(t, proto.Equal(laptop, &newLaptop))

	err = serializer.WriteProtobufToJSONFile(laptop, "new_laptop.json")
	require.NoError(t, err)
}
