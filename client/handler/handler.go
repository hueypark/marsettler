package handler

import (
	"encoding/binary"
	"fmt"
	"net"

	"github.com/hueypark/marsettler/client/ctx"
	"github.com/hueypark/marsettler/message"
)

// Handle handles message.
func Handle(conn net.Conn) error {
	if conn == nil {
		return fmt.Errorf("conn is nil")
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
		return fmt.Errorf("headsize is not same[expected: %d, got: %d]", size, readSize)
	}

	switch id {
	case message.MsgActorsPush:
		var actorsPush message.ActorsPush
		err := actorsPush.Unmarshal(body)
		if err != nil {
			return err
		}

		for _, actor := range actorsPush.Actors {
			ctx.World.NewActor(actor)
		}
	default:
		return fmt.Errorf("unhandled message id: %d", id)
	}

	return nil
}
