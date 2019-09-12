package message

type MsgID uint32

const (
	MsgActorPush MsgID = iota
)

// HeadSize represents size of head
const HeadSize = 8

type Msg interface {
	MsgID() MsgID
	MarshalTo(dAtA []byte) (int, error)
}

func (m *ActorsPush) MsgID() MsgID {
	return MsgActorPush
}
