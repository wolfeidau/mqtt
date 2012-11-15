// Implementation of MQTT V3.1.
// See http://public.dhe.ibm.com/software/dw/webservices/ws-mqtt/mqtt-v3r1.html
package mqtt

import (
	"errors"
	"io"
)

var (
	badMsgTypeError        = errors.New("mqtt: message type is invalid")
	badQosError            = errors.New("mqtt: QoS is invalid")
	badWillQosError        = errors.New("mqtt: will QoS is invalid")
	badLengthEncodingError = errors.New("mqtt: remaining length field exceeded maximum of 4 bytes")
	badReturnCodeError     = errors.New("mqtt: is invalid")
	dataExceedsPacketError = errors.New("mqtt: data exceeds packet length")
	msgTooLongError        = errors.New("mqtt: message is too long")
)

const (
	QosAtMostOnce = QosLevel(iota)
	QosAtLeastOnce
	QosExactlyOnce

	qosFirstInvalid
)

type QosLevel uint8

func (qos QosLevel) IsValid() bool {
	return qos < qosFirstInvalid
}

func (qos QosLevel) HasId() bool {
	return qos == QosAtLeastOnce || qos == QosExactlyOnce
}

const (
	RetCodeAccepted = ReturnCode(iota)
	RetCodeUnacceptableProtocolVersion
	RetCodeIdentifierRejected
	RetCodeServerUnavailable
	RetCodeBadUsernameOrPassword
	RetCodeNotAuthorized

	retCodeFirstInvalid
)

type ReturnCode uint8

func (rc ReturnCode) IsValid() bool {
	return rc >= RetCodeAccepted && rc < retCodeFirstInvalid
}

// DecodeOneMessage decodes one message from r.
func DecodeOneMessage(r io.Reader) (msg Message, err error) {
	var hdr Header
	var msgType MessageType
	var packetRemaining int32
	msgType, packetRemaining, err = hdr.Decode(r)

	msg, err = NewMessage(msgType)
	if err != nil {
		return
	}

	return msg, msg.Decode(r, hdr, packetRemaining)
}

// NewMessage creates an instance of a Message value for the given message
// type. An error is returned if msgType is invalid.
func NewMessage(msgType MessageType) (msg Message, err error) {
	switch msgType {
	case MsgConnect:
		msg = new(Connect)
	case MsgConnAck:
		msg = new(ConnAck)
	case MsgPublish:
		msg = new(Publish)
	case MsgPubAck:
		msg = new(PubAck)
	case MsgPubRec:
		msg = new(PubRec)
	case MsgPubRel:
		msg = new(PubRel)
	case MsgPubComp:
		msg = new(PubComp)
	case MsgSubscribe:
		msg = new(Subscribe)
	case MsgUnsubAck:
		msg = new(UnsubAck)
	case MsgSubAck:
		msg = new(SubAck)
	case MsgUnsubscribe:
		msg = new(Unsubscribe)
	case MsgPingReq:
		msg = new(PingReq)
	case MsgPingResp:
		msg = new(PingResp)
	case MsgDisconnect:
		msg = new(Disconnect)
	default:
		return nil, badMsgTypeError
	}

	return
}

// panicErr wraps an error that caused a problem that needs to bail out of the
// API, such that errors can be recovered and returned as errors from the
// public API.
type panicErr struct {
	err error
}

func (p panicErr) Error() string {
	return p.err.Error()
}

func raiseError(err error) {
	panic(panicErr{err})
}

// recoverError recovers any panic in flight and, iff it's an error from
// raiseError, will return the error. Otherwise re-raises the panic value.
// If no panic is in flight, it returns existingErr.
//
// This must be used in combination with a defer in all public API entry
// points where raiseError could be called.
func recoverError(existingErr error, recovered interface{}) error {
	if recovered != nil {
		if pErr, ok := recovered.(panicErr); ok {
			return pErr.err
		} else {
			panic(recovered)
		}
	}
	return existingErr
}
