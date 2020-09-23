// This file was generated from `./pkg/cmd/generate/generate_message_id.go`.

package message

import proto "github.com/golang/protobuf/proto"

type ID int32

const (
	ActorID ID = 0
	ActorMoveID ID = 1
	ActorMovesPushID ID = 2
	MoveStickRequestID ID = 3
	SignInRequestID ID = 4
	SignInResponseID ID = 5
	VectorID ID = 6
)

// Message represents message.
type Message interface {
	ID() ID
	proto.Message
}

func (m *Actor) ID() ID { return ActorID }
func (m *ActorMove) ID() ID { return ActorMoveID }
func (m *ActorMovesPush) ID() ID { return ActorMovesPushID }
func (m *MoveStickRequest) ID() ID { return MoveStickRequestID }
func (m *SignInRequest) ID() ID { return SignInRequestID }
func (m *SignInResponse) ID() ID { return SignInResponseID }
func (m *Vector) ID() ID { return VectorID }