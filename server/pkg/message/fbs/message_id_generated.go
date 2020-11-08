// This file was generated from `./pkg/cmd/generate/generate_message_id.go`.

package fbs

type ID int32

const (
	ActResponseID           ID = 0
	ActRequestID            ID = 1
	ActorID                 ID = 2
	ActorDisappearID        ID = 3
	ActorDisappearsPushID   ID = 4
	ActorMoveID             ID = 5
	ActorMovesPushID        ID = 6
	ActorsPushID            ID = 7
	MoveStickRequestID      ID = 8
	MoveToPositionRequestID ID = 9
	SignInRequestID         ID = 10
	SignInResponseID        ID = 11
	SkillUseRequestID       ID = 12
	SkillUseResponseID      ID = 13
	StatID                  ID = 14
	VectorID                ID = 15
)

func (rcv *SignInRequest) MessageID() ID { return SignInRequestID }
