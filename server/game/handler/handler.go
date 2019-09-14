package handler

import (
	"encoding/binary"
	"fmt"
	"net"

	"github.com/hueypark/marsettler/message"
	"github.com/hueypark/marsettler/server/game"
)

// Handle handle message
func Handle(userID int64, conn net.Conn) error {
	user := game.GetUser(userID)
	if user == nil {
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
	default:
		return fmt.Errorf("unhandled message id: %d", id)
	}

	return nil
}
