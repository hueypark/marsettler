package game

import (
	"encoding/binary"
	"log"
	"net"

	"github.com/hueypark/marsettler/game/message"
	"github.com/hueypark/marsettler/game/message/fbs"
)

// User represents user
type User struct {
	Conn net.Conn
}

// NewUser create new user
func NewUser(conn net.Conn) interface{} {
	return &User{conn}
}

// SendLoginResult sends login result
func (u User) SendLoginResult(id int64) {
	loginResult, size := message.MakeLoginResult(id)

	u.sendHead(fbs.LoginResultID, size)
	u.sendBody(loginResult)
}

func (u User) sendHead(id fbs.MessageID, size int) {
	head := make([]byte, fbs.HeadSize)

	binary.LittleEndian.PutUint32(head[0:], uint32(id))
	binary.LittleEndian.PutUint32(head[4:], uint32(size))

	_, err := u.Conn.Write(head)
	if err != nil {
		log.Println(err)
	}
}

func (u User) sendBody(bytes []byte) {
	_, err := u.Conn.Write(bytes)
	if err != nil {
		log.Println(err)
	}
}
