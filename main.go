package main

import (
	"encoding/binary"
	"log"
	"net"

	"github.com/hueypark/marsettler/message"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}

		for {
			head := make([]byte, 8)
			body := message.MakeActor(291452, 111210.0, 312312.0)

			id := 197
			size := len(body)

			binary.LittleEndian.PutUint32(head[0:], uint32(id))
			binary.LittleEndian.PutUint32(head[4:], uint32(size))

			_, err = conn.Write(head)
			if err != nil {
				log.Println(err)
				break
			}
			_, _ = conn.Write(body)
			if err != nil {
				log.Println(err)
				break
			}
		}
	}
}
