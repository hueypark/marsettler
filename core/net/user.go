package net

type User interface {
	ID() int64
	OnCreated()
	OnMessage(bytes []byte)
	OnClosed()
}
