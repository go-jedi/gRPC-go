package serializer

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ProtobufToJSON(message proto.Message) ([]byte, error) {
	marshalerOpts := protojson.MarshalOptions{
		UseEnumNumbers:  true,
		Indent:          " ",
		EmitUnpopulated: true,
	}
	marshaler, err := marshalerOpts.Marshal(message)
	if err != nil {
		return []byte{}, err
	}
	return marshaler, nil
}
