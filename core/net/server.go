package net

import (
	"log"
	"net"
)

// Server represent server
type Server struct {
	newConns  chan net.Conn
	deadConns chan net.Conn
	newUser   func(conn net.Conn) interface{}
	handler   func(user interface{}) error
}

// NewServer create server
func NewServer(
	newUser func(conn net.Conn) interface{},
	handler func(user interface{}) error,
) *Server {
	server := &Server{
		make(chan net.Conn),
		make(chan net.Conn),
		newUser,
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
	user := server.newUser(conn)

	for {
		err := server.read(user)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func (server *Server) read(user interface{}) error {
	return server.handler(user)
}
