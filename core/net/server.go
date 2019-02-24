package net

import (
	"log"
	"net"

	"github.com/hueypark/marsettler/core/id_generator"
)

// Server represent server
type Server struct {
	newConns  chan net.Conn
	deadConns chan net.Conn
	onAccept  func(userID int64, conn net.Conn)
	onClose   func(userID int64)
	handler   func(userID int64, conn net.Conn) error
}

// NewServer create server
func NewServer(
	onAccept func(userID int64, conn net.Conn),
	onClose func(userID int64),
	handler func(userID int64, conn net.Conn) error,
) *Server {
	server := &Server{
		make(chan net.Conn),
		make(chan net.Conn),
		onAccept,
		onClose,
		handler,
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
			go func() {
				server.handle(conn)
			}()
		case deadConn := <-server.deadConns:
			err := deadConn.Close()
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func (server *Server) handle(conn net.Conn) {
	id := id_generator.Generate()
	server.onAccept(id, conn)

	for {
		err := server.handler(id, conn)
		if err != nil {
			log.Println(err)
			break
		}
	}

	server.onClose(id)
	server.deadConns <- conn
}
