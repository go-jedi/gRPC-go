package serializer

import (
	"testing"

	"github.com/rob-bender/go-test/pb"
	"github.com/rob-bender/go-test/sample"
	"github.com/stretchr/testify/require"
)

func Test_WriteProtobufToJSONFile(t *testing.T) {
	t.Parallel()
	var binaryFile string = "../tmp/laptop.bin"
	var jsonFile string = "../tmp/laptop.json"

	lapton1 := sample.NewLaptor()
	err := WriteProtobufToBinaryFile(lapton1, binaryFile)
	require.NoError(t, err)

	err = WriteProtobufToJSONFile(lapton1, jsonFile)
	require.NoError(t, err)
}

func Test_WriteProtobufToBinaryFile(t *testing.T) {
	t.Parallel()

	var binaryFile string = "../tmp/laptop.bin"
	lapton1 := sample.NewLaptor()
	err := WriteProtobufToBinaryFile(lapton1, binaryFile)
	require.NoError(t, err)
}

func Test_ReadrProtobufFromBinaryFile(t *testing.T) {
	t.Parallel()

	var binaryFile string = "../tmp/laptop.bin"
	lapton2 := &pb.Laptor{}
	err := ReadrProtobufFromBinaryFile(binaryFile, lapton2)
	require.NoError(t, err)
}
