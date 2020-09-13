// This file was generated from `./pkg/cmd/generate/generate_message_id.go`.

package message

import proto "github.com/golang/protobuf/proto"

type ID int32

const (
	PingID ID = 0
	PongID ID = 1
)

// Message represents message.
type Message interface {
	ID() ID
	proto.Message
}

func (m *Ping) ID() ID { return PingID }
func (m *Pong) ID() ID { return PongID }
