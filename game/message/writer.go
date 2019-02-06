package message

import (
	"encoding/binary"
	"log"
	"net"

	"github.com/hueypark/marsettler/game/message/fbs"
)

// WriteHead sends message head.
func WriteHead(conn net.Conn, id fbs.MessageID, size int) {
	head := make([]byte, fbs.HeadSize)

	binary.LittleEndian.PutUint32(head[0:], uint32(id))
	binary.LittleEndian.PutUint32(head[4:], uint32(size))

	_, err := conn.Write(head)
	if err != nil {
		log.Println(err)
	}
}

// WriteBody sends message body.
func WriteBody(conn net.Conn, bytes []byte) {
	_, err := conn.Write(bytes)
	if err != nil {
		log.Println(err)
	}
}
