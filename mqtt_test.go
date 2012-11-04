package mqtt

import (
	"bytes"
	"reflect"
	"testing"

	gbt "github.com/huin/gobinarytest"
)

var bitCnt = uint32(0)

func Test(t *testing.T) {
	mqtt := Mqtt{
		Header: Header{MessageType: CONNECT},
		ProtocolName: "MQIsdp",
		ProtocolVersion: 3,
		ConnectFlags: ConnectFlags{
			UsernameFlag: true,
			PasswordFlag: true,
			WillRetain: false,
			WillQos: 1,
			WillFlag: true,
			CleanSession: true,
		},
		KeepAliveTimer: 10,
		ClientId: "xixihaha",
		WillTopic: "topic",
		WillMessage: "message",
		Username: "name",
		Password: "pwd",
	}

	expected := gbt.InOrder{
		gbt.Named{"Header byte", gbt.Literal{0x10}},
		gbt.Named{"Remaining length", gbt.Literal{12 + 5*2 + 8 + 5 + 7 + 4 + 3}},

		// Extended headers for CONNECT:
		gbt.Named{"Protocol name", gbt.InOrder{gbt.Literal{0x00, 0x06}, gbt.Literal("MQIsdp")}},
		gbt.Named{
			"Extended headers for CONNECT",
			gbt.Literal{
				0x03, // Protocol version number
				0xce, // Connect flags
				0x00, 0x0a, // Keep alive timer
			},
		},

		// CONNECT payload:
		gbt.Named{"Client identifier", gbt.InOrder{gbt.Literal{0x00, 0x08}, gbt.Literal("xixihaha")}},
		gbt.Named{"Will topic", gbt.InOrder{gbt.Literal{0x00, 0x05}, gbt.Literal("topic")}},
		gbt.Named{"Will message", gbt.InOrder{gbt.Literal{0x00, 0x07}, gbt.Literal("message")}},
		gbt.Named{"Username", gbt.InOrder{gbt.Literal{0x00, 0x04}, gbt.Literal("name")}},
		gbt.Named{"Password", gbt.InOrder{gbt.Literal{0x00, 0x03}, gbt.Literal("pwd")}},
	}

	if encoded, err := Encode(&mqtt); err != nil {
		t.Errorf("Unexpected error during encoding: %v", err)
	} else if err = gbt.Matches(expected, encoded); err != nil {
		t.Errorf("Unexpected encoding output: %v", err)
	}

	encodedBuf := new(bytes.Buffer)
	expected.Write(encodedBuf)

	if decodedMqtt, err := Decode(encodedBuf.Bytes()); err != nil {
		t.Errorf("Unexpected error during decoding: %v", err)
	} else if !reflect.DeepEqual(&mqtt, decodedMqtt) {
		t.Errorf("Decoded value mismatch\n     got = %#v\nexpected = %#v", *decodedMqtt, mqtt)
	}
}
