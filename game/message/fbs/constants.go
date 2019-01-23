package fbs

// MessageID represents message id
type MessageID uint32

// MessageID
const (
	ActorID MessageID = iota
	LoginID
	LoginResultID
)

// HeadSize represents size of head
const HeadSize = 8
