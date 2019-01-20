package net

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Server struct {
	newUser func(*websocket.Conn) User
}

func NewServer(newUser func(*websocket.Conn) User) *Server {
	server := &Server{
		newUser,
	}

	return server
}

func (server Server) ListenAndServe() {
	http.HandleFunc(
		"/",
		func(response http.ResponseWriter, request *http.Request) { server.handler(response, request) })
	http.ListenAndServe("localhost:8080", nil)
}

func (server *Server) handler(response http.ResponseWriter, request *http.Request) {
	var upgrader websocket.Upgrader
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := upgrader.Upgrade(response, request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	user := server.newUser(conn)
	user.OnCreated()

	defer func() {
		user.OnClosed()
		conn.Close()
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Read failed. [err: %v]", err)
			break
		}

		user.OnMessage(message)
	}
}
