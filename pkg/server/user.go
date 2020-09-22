package server

import (
	"github.com/hueypark/marsettler/pkg/global"
	"github.com/hueypark/marsettler/pkg/internal/net"
	"github.com/hueypark/marsettler/pkg/message"
	"github.com/hueypark/marsettler/pkg/server/game"
)

// User represents user.
type User struct {
	id    int64
	conn  *net.Conn
	actor *game.Actor
}

// NewUser creates new user.
func NewUser(conn *net.Conn) *User {
	u := &User{
		id:   global.IdGenerator.Generate().Int64(),
		conn: conn,
	}

	return u
}

// Actor returns actor.
func (u *User) Actor() *game.Actor {
	return u.actor
}

// ID returns id.
func (u *User) ID() int64 {
	return u.id
}

// SetActor sets actor.
func (u *User) SetActor(actor *game.Actor) {
	u.actor = actor
}

// Write writes message.
func (u *User) Write(message message.Message) error {
	return u.conn.Write(message)
}
