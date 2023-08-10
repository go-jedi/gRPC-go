package serializer

import (
	"fmt"
	"os"

	"google.golang.org/protobuf/proto"
)

func WriteProtobufToJSONFile(message proto.Message, filename string) error {
	dataByte, err := ProtobufToJSON(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to JSON: %w", err)
	}
	err = os.WriteFile(filename, dataByte, 0644)
	if err != nil {
		return fmt.Errorf("cannot write JSON data to file: %w", err)
	}
	return nil
}

func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to binary: %w", err)
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("cannot write binary data to file: %w", err)
	}
	return nil
}

func ReadrProtobufFromBinaryFile(filename string, message proto.Message) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cannot read binary data from file: %w", err)
	}
	err = proto.Unmarshal(file, message)
	if err != nil {
		return fmt.Errorf("cannot unmarshal binary to proto message: %w", err)
	}
	return nil
}
