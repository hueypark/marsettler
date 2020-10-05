// This file was generated from `./pkg/cmd/generate/generate_message_id.go`.

package message

import proto "github.com/golang/protobuf/proto"

type ID int32

const (
	ActResponseID ID = 0
	ActRequestID ID = 1
	ActorID ID = 2
	ActorDisappearID ID = 3
	ActorDisappearsPushID ID = 4
	ActorMoveID ID = 5
	ActorMovesPushID ID = 6
	ActorsPushID ID = 7
	MoveStickRequestID ID = 8
	MoveToPositionRequestID ID = 9
	SignInRequestID ID = 10
	SignInResponseID ID = 11
	VectorID ID = 12
)

// Message represents message.
type Message interface {
	ID() ID
	proto.Message
}

func (m *ActResponse) ID() ID { return ActResponseID }
func (m *ActRequest) ID() ID { return ActRequestID }
func (m *Actor) ID() ID { return ActorID }
func (m *ActorDisappear) ID() ID { return ActorDisappearID }
func (m *ActorDisappearsPush) ID() ID { return ActorDisappearsPushID }
func (m *ActorMove) ID() ID { return ActorMoveID }
func (m *ActorMovesPush) ID() ID { return ActorMovesPushID }
func (m *ActorsPush) ID() ID { return ActorsPushID }
func (m *MoveStickRequest) ID() ID { return MoveStickRequestID }
func (m *MoveToPositionRequest) ID() ID { return MoveToPositionRequestID }
func (m *SignInRequest) ID() ID { return SignInRequestID }
func (m *SignInResponse) ID() ID { return SignInResponseID }
func (m *Vector) ID() ID { return VectorID }
