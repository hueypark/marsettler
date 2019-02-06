package message

import (
	"encoding/binary"
	"fmt"

	"github.com/hueypark/marsettler/game/message/fbs"
)

// ReadMessage reads message from conn.
func ReadMessage(client client) (id fbs.MessageID, body []byte, err error) {
	conn := client.Conn()
	if conn == nil {
		return id, body, fmt.Errorf("conn is null")
	}

	head := make([]byte, fbs.HeadSize)
	read, err := conn.Read(head)
	if err != nil {
		return id, body, err
	}
	if read != fbs.HeadSize {
		return id, body, fmt.Errorf("headsize is not same[expected: %d, got: %d]", fbs.HeadSize, read)
	}

	id = (fbs.MessageID)(binary.LittleEndian.Uint32(head[0:]))
	size := (int)(binary.LittleEndian.Uint32(head[4:]))

	body = make([]byte, size)
	read, err = conn.Read(body)
	if err != nil {
		return id, body, err
	}
	if read != size {
		return id, body, fmt.Errorf("headsize is not same[expected: %d, got: %d]", size, read)
	}

	return id, body, nil
}
