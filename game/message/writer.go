package message

import (
	"encoding/binary"
	"fmt"

	"github.com/hueypark/marsettler/game/message/fbs"
)

// WriteHead sends message head.
func WriteHead(client client, id fbs.MessageID, size int) error {
	conn := client.Conn()
	if conn == nil {
		return fmt.Errorf("conn is null")
	}

	head := make([]byte, fbs.HeadSize)

	binary.LittleEndian.PutUint32(head[0:], uint32(id))
	binary.LittleEndian.PutUint32(head[4:], uint32(size))

	_, err := conn.Write(head)
	return err
}

// WriteBody sends message body.
func WriteBody(client client, bytes []byte) error {
	conn := client.Conn()
	if conn == nil {
		return fmt.Errorf("conn is null")
	}

	_, err := client.Conn().Write(bytes)
	return err
}
