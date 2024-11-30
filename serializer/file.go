package serializer

import (
	"fmt"
	"os"

	"google.golang.org/protobuf/proto"
)

func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("can not marshal proto message to binary : %w", err)
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}
	_, err = file.Write(data)

	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}

func ReadProtobufFromBinaryFile(filename string, message proto.Message) error {
	data, err := os.ReadFile(filename)

	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	err = proto.Unmarshal(data, message)

	if err != nil {
		return fmt.Errorf("error unmarshaling file %w", err)
	}

	return nil
}

func WriteProtobufToJSONFile(message proto.Message, filename string) error {
	json, err := ProtobufToJSON(message)

	if err != nil {
		return fmt.Errorf("error converting protobuf message to Json: %w", err)
	}

	file, err := os.Create(filename)

	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}

	_, err = file.Write(json)

	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}
