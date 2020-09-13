package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/hueypark/marsettler/pkg/message"
	"github.com/hueypark/marsettler/pkg/shared"
)

// Server is the marsettler server.
type Server struct {
	gin      *gin.Engine
	upgrader websocket.Upgrader
}

// NewServer creates new server.
func NewServer() *Server {
	s := &Server{
		gin:      gin.Default(),
		upgrader: websocket.Upgrader{},
	}

	s.gin.GET(
		"/ws",
		func(c *gin.Context) {
			err := upgrade(c.Writer, c.Request)
			if err != nil {
				log.Println(err)
				return
			}
		})

	return s
}

// Run runs server.
func (s *Server) Run() error {
	return s.gin.Run()
}

func upgrade(w http.ResponseWriter, r *http.Request) error {
	var upgrader websocket.Upgrader

	websocketConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}

	handlers := shared.HandlerFuncs{
		message.PingID: func(conn *shared.Conn, m *message.Ping) error {
			pong := &message.Pong{}
			return conn.Write(pong)
		},
	}

	conn, err := shared.NewConn(websocketConn, handlers)
	if err != nil {
		return err
	}

	go conn.Run()

	for {
		conn.Consume()
	}

	//defer conn.Close()
	//
	//for {
	//	_, message, err := conn.ReadMessage()
	//	if err != nil {
	//		log.Println(err)
	//		break
	//	}
	//
	//	log.Println(string(message))
	//
	//	err = conn.WriteMessage(websocket.BinaryMessage, message)
	//	if err != nil {
	//		log.Println(err)
	//		break
	//	}
	//}

	return nil
}
