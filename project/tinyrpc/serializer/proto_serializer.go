package serializer

import (
	"errors"

	"google.golang.org/protobuf/proto"
)

var NotImplementProtoMessageError = errors.New("param does not implement proto.Message")

var Proto = ProtoSerializer{}

// ProtoSerializer implements the Serializer interface
type ProtoSerializer struct{}

// Marshal .
func (_ ProtoSerializer) Marshal(message interface{}) ([]byte, error) {
	if message == nil {
		return []byte{}, nil
	}
	body, ok := message.(proto.Message)
	if !ok {
		return nil, NotImplementProtoMessageError
	}
	return proto.Marshal(body)
}

// Unmarshal .
func (_ ProtoSerializer) Unmarshal(data []byte, message interface{}) error {
	if message == nil {
		return nil
	}
	body, ok := message.(proto.Message)
	if !ok {
		return NotImplementProtoMessageError
	}
	return proto.Unmarshal(data, body)
}
