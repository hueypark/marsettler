package message

import proto "github.com/golang/protobuf/proto"

type ID int32

const (
	PingID ID = iota
	PongID
)

// Message represents message.
type Message interface {
	ID() ID
	proto.Message
}

// ID represents message ID.
func (m *Ping) ID() ID {
	return PingID
}

// ID represents messageID.
func (m *Pong) ID() ID {
	return PongID

}
