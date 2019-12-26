package user

import (
	"encoding/binary"
	"log"
	"net"
	"sync"

	"github.com/gogo/protobuf/proto"
	"github.com/hueypark/marsettler/message"
	"github.com/hueypark/marsettler/pkg/ctx"
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
	mux.RLock()
	defer mux.RUnlock()
	return users[userID]
}

// OnAccept handles net.Conn's accept event.
func OnAccept(userID int64, conn net.Conn) {
	mux.Lock()
	defer mux.Unlock()

	user := &User{
		conn: conn,
	}

	users[userID] = user

	ctx.World.AddListener(user)
}

// OnClose handles net.Conn's close event.
func OnClose(userID int64) {
	mux.Lock()
	defer mux.Unlock()

	delete(users, userID)
}

// Send sends message.
func (u *User) Send(msg message.Msg) {
	if u.conn == nil {
		log.Println("conn is nil")
		return
	}

	id := msg.MsgID()
	size := msg.Size()
	buffer, err := proto.Marshal(msg)
	if err != nil {
		log.Println(err)
	}

	head := make([]byte, message.HeadSize)
	binary.LittleEndian.PutUint32(head[0:], uint32(id))
	binary.LittleEndian.PutUint32(head[4:], uint32(size))

	_, err = u.conn.Write(head)
	if err != nil {
		log.Println(err)
	}

	_, err = u.conn.Write(buffer)
	if err != nil {
		log.Println(err)
	}
}
