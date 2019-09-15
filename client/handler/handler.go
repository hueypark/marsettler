package handler

import (
	"encoding/binary"
	"fmt"
	"net"

	"github.com/gogo/protobuf/proto"
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
		return fmt.Errorf("bodySize is not same[expected: %d, got: %d]", size, readSize)
	}

	switch id {
	case message.MsgActors:
		actors := &message.Actors{}
		err := proto.Unmarshal(body, actors)
		if err != nil {
			return err
		}
		return handleActors(actors)
	case message.MsgWorld:
		world := &message.World{}
		err := proto.Unmarshal(body, world)
		if err != nil {
			return err
		}
		return handleWorld(world)
	default:
		return fmt.Errorf("unhandled message id: %d", id)
	}
}
