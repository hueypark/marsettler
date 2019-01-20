package net

import (
	"encoding/binary"
	"log"
	"net"

	"github.com/hueypark/marsettler/message"
)

// Server represent server
type Server struct {
	newConns  chan net.Conn
	deadConns chan net.Conn
	conns     map[net.Conn]bool
}

// NewServer create server
func NewServer() *Server {
	server := &Server{
		make(chan net.Conn),
		make(chan net.Conn),
		make(map[net.Conn]bool),
	}

	return server
}

// Listen open port and listen to connections
func (server Server) Listen(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		err := listener.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Println(err)
			}

			server.newConns <- conn
		}
	}()

	for {
		select {
		case conn := <-server.newConns:
			server.conns[conn] = true
			go func() {
				server.handle(conn)
			}()
		case deadConn := <-server.deadConns:
			err := deadConn.Close()
			if err != nil {
				log.Println(err)
			}
			delete(server.conns, deadConn)
		}
	}
}

func (server *Server) handle(conn net.Conn) {
	for {
		head := make([]byte, 8)
		body := message.MakeActor(291452, 111210.0, 312312.0)

		id := 197
		size := len(body)

		binary.LittleEndian.PutUint32(head[0:], uint32(id))
		binary.LittleEndian.PutUint32(head[4:], uint32(size))

		_, err := conn.Write(head)
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
