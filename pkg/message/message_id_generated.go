// This file was generated from `./pkg/cmd/generate/generate_message_id.go`.

package message

import proto "github.com/golang/protobuf/proto"

type ID int32

const (
	ActorID ID = 0
	ActorMoveID ID = 1
	ActorMovesPushID ID = 2
	MoveStickID ID = 3
	PingID ID = 4
	PongID ID = 5
	SignInID ID = 6
	SignInResponseID ID = 7
	VectorID ID = 8
)

// Message represents message.
type Message interface {
	ID() ID
	proto.Message
}

func (m *Actor) ID() ID { return ActorID }
func (m *ActorMove) ID() ID { return ActorMoveID }
func (m *ActorMovesPush) ID() ID { return ActorMovesPushID }
func (m *MoveStick) ID() ID { return MoveStickID }
func (m *Ping) ID() ID { return PingID }
func (m *Pong) ID() ID { return PongID }
func (m *SignIn) ID() ID { return SignInID }
func (m *SignInResponse) ID() ID { return SignInResponseID }
func (m *Vector) ID() ID { return VectorID }
