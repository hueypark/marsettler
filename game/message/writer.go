package message

import (
	"encoding/binary"
	"log"
	"net"

	"github.com/hueypark/marsettler/game/message/fbs"
)

// Write writes message.
func Write(conn net.Conn, id fbs.MessageID, bytes []byte) {
	size := uint32(len(bytes))

	head := make([]byte, fbs.HeadSize)

	binary.LittleEndian.PutUint32(head[0:], uint32(id))
	binary.LittleEndian.PutUint32(head[4:], size)

	_, err := conn.Write(head)
	if err != nil {
		log.Println(err)
	}

	_, err = conn.Write(bytes)
	if err != nil {
		log.Println(err)
	}
}
