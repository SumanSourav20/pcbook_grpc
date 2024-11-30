package serializer

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ProtobufToJSON(message proto.Message) ([]byte, error) {
	options := protojson.MarshalOptions{
		UseEnumNumbers:    false,
		EmitDefaultValues: true,
		Indent:            "	",
	}

	return options.Marshal(message)
}
