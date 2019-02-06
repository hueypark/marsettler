package game

import (
	"net"

	"github.com/hueypark/marsettler/game/message"
	"github.com/hueypark/marsettler/game/message/fbs"
)

// User represents user
type User struct {
	conn net.Conn
}

// NewUser create new user
func NewUser(conn net.Conn) interface{} {
	return &User{conn}
}

// Conn returns network connection.
func (u *User) Conn() net.Conn {
	return u.conn
}

// SendLoginResult sends login result
func (u *User) SendLoginResult(id int64) {
	loginResult, size := message.MakeLoginResultMessage(id)

	message.WriteHead(u, fbs.LoginResultID, size)
	message.WriteBody(u, loginResult)
}
