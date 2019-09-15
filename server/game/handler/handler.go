package handler

import (
	"encoding/binary"
	"fmt"
	"net"

	"github.com/gogo/protobuf/proto"
	"github.com/hueypark/marsettler/message"
	"github.com/hueypark/marsettler/server/user"
)

// Handle handle message
func Handle(userID int64, conn net.Conn) error {
	u := user.GetUser(userID)
	if u == nil {
		return fmt.Errorf("user is nil id: %d", userID)
	}

	head := make([]byte, message.HeadSize)
	readSize, err := conn.Read(head)
	if err != nil {
		return err
	}
	if readSize != message.HeadSize {
		return fmt.Errorf("headsize is not same[expected: %d, got: %d]", message.HeadSize, readSize)
	}

	id := (message.MsgID)(binary.LittleEndian.Uint32(head[0:]))
	size := (int)(binary.LittleEndian.Uint32(head[4:]))

	body := make([]byte, size)
	readSize, err = conn.Read(body)
	if err != nil {
		return err
	}
	if readSize != size {
		return fmt.Errorf("bodySize is not same[expected: %d, got: %d]", size, readSize)
	}

	switch id {
	case message.MsgActorCreate:
		actorCreate := &message.ActorCreate{}
		err := proto.Unmarshal(body, actorCreate)
		if err != nil {
			return err
		}
		return handleActorCreate(actorCreate)
	default:
		return fmt.Errorf("unhandled message id: %d", id)
	}
}
