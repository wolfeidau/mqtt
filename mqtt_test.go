package mqtt

import (
	"fmt"
	"testing"
)

var bitCnt = uint32(0)

func Test(t *testing.T) {
	mqtt := initTest()
	fmt.Println("------ Origin MQTT Object")
	printMqtt(mqtt)
	fmt.Println("------ Encode To Binary")
	bits, _ := Encode(mqtt)
	printBytes(bits)
	fmt.Println("------ Decode To Object")
	newMqtt, _ := Decode(bits)
	printMqtt(newMqtt)
}

func initTest() *Mqtt {
	return &Mqtt{
		Header: &Header{MessageType: CONNECT},
		ProtocolName: "MQIsdp",
		ProtocolVersion: 3,
		ConnectFlags: &ConnectFlags{
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
}

func printByte(b byte) {
	bitCnt += 1
	out := make([]uint8, 8)
	val := uint8(b)
	for i := 1; val > 0; i += 1 {
		foo := val % 2
		val = (val - foo) / 2
		out[8-i] = foo
	}
	fmt.Println(bitCnt, out)
}

func printBytes(b []byte) {
	for i := 0; i < len(b); i += 1 {
		printByte(b[i])
	}
}

func printMqtt(mqtt *Mqtt) {
	fmt.Printf("MQTT = %+v\n", *mqtt)
	fmt.Printf("Header = %+v\n", *mqtt.Header)
	if mqtt.ConnectFlags != nil {
		fmt.Printf("ConnectFlags = %+v\n", *mqtt.ConnectFlags)
	}
}
