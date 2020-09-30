// This file was generated from `./pkg/cmd/generate/generate_message_id.go`.

package message

import proto "github.com/golang/protobuf/proto"

type ID int32

const (
	ActResponseID ID = 0
	ActRequestID ID = 1
	ActorID ID = 2
	ActorMoveID ID = 3
	ActorMovesPushID ID = 4
	ActorsPushID ID = 5
	MoveStickRequestID ID = 6
	SignInRequestID ID = 7
	SignInResponseID ID = 8
	VectorID ID = 9
)

// Message represents message.
type Message interface {
	ID() ID
	proto.Message
}

func (m *ActResponse) ID() ID { return ActResponseID }
func (m *ActRequest) ID() ID { return ActRequestID }
func (m *Actor) ID() ID { return ActorID }
func (m *ActorMove) ID() ID { return ActorMoveID }
func (m *ActorMovesPush) ID() ID { return ActorMovesPushID }
func (m *ActorsPush) ID() ID { return ActorsPushID }
func (m *MoveStickRequest) ID() ID { return MoveStickRequestID }
func (m *SignInRequest) ID() ID { return SignInRequestID }
func (m *SignInResponse) ID() ID { return SignInResponseID }
func (m *Vector) ID() ID { return VectorID }
