package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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
			conn, err := s.upgrader.Upgrade(c.Writer, c.Request, nil)
			if err != nil {
				log.Println(err)
				return
			}
			defer conn.Close()

			for {
				_, message, err := conn.ReadMessage()
				if err != nil {
					log.Println(err)
					break
				}

				log.Println(string(message))

				err = conn.WriteMessage(websocket.BinaryMessage, message)
				if err != nil {
					log.Println(err)
					break
				}
			}
		})

	return s
}

// Run runs server.
func (s *Server) Run() error {
	return s.gin.Run()
}
