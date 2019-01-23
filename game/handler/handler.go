package handler

import (
	"encoding/binary"
	"fmt"

	"github.com/hueypark/marsettler/game"

	"github.com/hueypark/marsettler/game/message"
	"github.com/hueypark/marsettler/game/message/fbs"
)

// Handle handle message
func Handle(iUser interface{}) error {
	user := iUser.(game.User)

	head := make([]byte, fbs.HeadSize)
	read, err := user.Conn.Read(head)
	if err != nil {
		return err
	}
	if read != fbs.HeadSize {
		return fmt.Errorf("headsize is not same[expected: %d, got: %d]", fbs.HeadSize, read)
	}

	id := binary.LittleEndian.Uint32(head[0:])
	size := (int)(binary.LittleEndian.Uint32(head[4:]))

	body := make([]byte, size)
	read, err = user.Conn.Read(body)
	if err != nil {
		return err
	}
	if read != size {
		return fmt.Errorf("headsize is not same[expected: %d, got: %d]", size, read)
	}

	messageID := fbs.MessageID(id)
	switch messageID {
	case fbs.LoginID:
		login := message.MakeLogin(body)
		return handleLogin(user, login)
	}

	return fmt.Errorf("undefined message id: %d", id)
}
