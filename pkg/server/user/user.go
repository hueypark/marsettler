package user

import (
	"github.com/hueypark/marsettler/pkg/internal/net"
	"github.com/hueypark/marsettler/pkg/message"
	"github.com/hueypark/marsettler/pkg/server/game"
)

// User represents user.
type User struct {
	id        int64
	conn      *net.Conn
	closeChan chan bool
	closed    bool
	actor     *game.Actor
}

// New creates new user.
func New(conn *net.Conn) *User {
	u := &User{
		conn: conn,
	}

	return u
}

// Actor returns actor.
func (u *User) Actor() *game.Actor {
	return u.actor
}

// Close closes user.
func (u *User) Close() {
	u.conn.Close()
	u.closed = true
}

// Consume consumes connection.
func (u *User) Consume() error {
	return u.conn.Consume()
}

// Expired returns true when the user expired.
func (u *User) Expired() bool {
	// TODO(jaewan): Expire state with a timeout. We can not verify this with a network event.
	return u.closed
}

// ID returns id.
func (u *User) ID() int64 {
	return u.id
}

// SetActor sets actor.
func (u *User) SetActor(actor *game.Actor) {
	u.actor = actor

	u.actor.SetWriter(
		func(message message.Message) error {
			return u.Write(message)
		})
}

// SetID sets id.
func (u *User) SetID(id int64) {
	u.id = id
}

// Write writes message.
func (u *User) Write(message message.Message) error {
	return u.conn.Write(message)
}
