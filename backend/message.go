package backend

import (
	"fmt"
	"log"

	"github.com/pierrec/lz4/v4"
	"google.golang.org/protobuf/proto"
)

func DecodeProto(message []byte) *Response {
	var dataframe Frame

	err := proto.Unmarshal(message, &dataframe)

	if err != nil {
		log.Panicln("parse dataframe error:", err)
	}

	truePayload := make([]byte, len(dataframe.Payload)*10)

	size, err := lz4.UncompressBlock(dataframe.Payload, truePayload)
	if err != nil {

		for _, b := range truePayload {
			fmt.Printf("%02x ", b)
		}
		fmt.Println()

		panic(err)
	}

	var response Response

	err = proto.Unmarshal(truePayload[:size], &response)

	if err != nil {
		log.Panicln("parse response payload error:", err)
	}

	return &response
}
