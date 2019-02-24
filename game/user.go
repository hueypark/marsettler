package game

import (
	"net"
	"sync"

	"github.com/hueypark/marsettler/game/message"
	"github.com/hueypark/marsettler/game/message/fbs"
)

var (
	mux   sync.RWMutex
	users = map[int64]*User{}
)

// User represents user
type User struct {
	conn net.Conn
}

// GetUser returns user.
func GetUser(userID int64) *User {
	return users[userID]
}

// OnAccept handles net.Conn's accept event.
func OnAccept(userID int64, conn net.Conn) {
	mux.Lock()
	defer mux.Unlock()

	user := &User{conn}

	users[userID] = user
}

// OnClose handles net.Conn's close event.
func OnClose(userID int64) {
	mux.Lock()
	defer mux.Unlock()

	delete(users, userID)
}

// SendLoginResult sends login result
func (u *User) SendLoginResult(id int64) {
	loginResult := message.MakeLoginResultMessage(id)

	message.Write(u.conn, fbs.LoginResultID, loginResult)
}
