package backend

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
)

func (akun *Account) CreateFrame() *Frame {
	akun.SequenceId += 1

	return &Frame{
		Seqid:       akun.SequenceId,
		Service:     10002,
		PayloadType: "pb",
		Method:      1,
		// Logid:       10001,
	}
}

func (akun *Account) CreateV2Init() []byte {

	data := &Request{
		Headers:        map[string]string{},
		AuthType:       2,
		DevicePlatform: "web",
		InboxType:      0,
		BuildNumber:    "12c929a:master",
		SdkVersion:     "0.3.8",
		Cmd:            203,
		Body: &RequestBody{
			MessagesPerUserInitV2Body: &MessagesPerUserInitV2RequestBody{
				Cursor: 0,
			},
		},
		Refer:      3,
		Token:      akun.Token,
		DeviceId:   akun.DeviceId,
		SequenceId: 10001,
	}

	payload, err := proto.Marshal(data)

	frame := akun.CreateFrame()
	frame.Payload = payload

	if err != nil {
		panic("gagal create request websocket")
	}

	hasil, _ := proto.Marshal(frame)
	return hasil
}

func (akun *Account) createPing() []byte {

	data := &Request{
		Headers: map[string]string{},
		Cmd:     200,
		Refer:   3,
		Body: &RequestBody{
			MessagesPerUserBody: &MessagesPerUserRequestBody{
				Cursor: akun.CursorChat,
				Limit:  50,
			},
		},
		Token:          akun.Token,
		DeviceId:       akun.DeviceId,
		AuthType:       2,
		DevicePlatform: "web",
		InboxType:      0,
		BuildNumber:    "12c929a:master",
		SdkVersion:     "0.3.8",
	}

	payload, err := proto.Marshal(data)

	frame := &Frame{
		Seqid:       10001,
		Logid:       10001,
		Service:     10002,
		PayloadType: "pb",
		Method:      1,
		Payload:     payload,
	}

	if err != nil {
		panic("gagal create request websocket")
	}

	hasil, _ := proto.Marshal(frame)

	log.Println("creating ping", fmt.Sprintf("%x", hasil))

	return hasil
}
