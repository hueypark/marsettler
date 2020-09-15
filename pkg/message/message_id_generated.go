// This file was generated from `./pkg/cmd/generate/generate_message_id.go`.

package message

import proto "github.com/golang/protobuf/proto"

type ID int32

const (
	ActorID ID = 0
	PingID ID = 1
	PongID ID = 2
	SignInID ID = 3
	SignInResponseID ID = 4
	VectorID ID = 5
)

// Message represents message.
type Message interface {
	ID() ID
	proto.Message
}

func (m *Actor) ID() ID { return ActorID }
func (m *Ping) ID() ID { return PingID }
func (m *Pong) ID() ID { return PongID }
func (m *SignIn) ID() ID { return SignInID }
func (m *SignInResponse) ID() ID { return SignInResponseID }
func (m *Vector) ID() ID { return VectorID }
