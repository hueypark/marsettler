package message

type MsgID uint32

const (
	MsgActors MsgID = iota
	MsgWorld
)

// HeadSize represents size of head
const HeadSize = 8

type Msg interface {
	MsgID() MsgID
	Size() (n int)
	Reset()
	String() string
	ProtoMessage()
}

func (m *Actors) MsgID() MsgID {
	return MsgActors
}

func (m *World) MsgID() MsgID {
	return MsgWorld
}
