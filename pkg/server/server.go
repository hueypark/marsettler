package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/hueypark/marsettler/pkg/internal/net"
	"github.com/hueypark/marsettler/pkg/message"
	"github.com/hueypark/marsettler/pkg/server/game"
)

// Server is the client server.
type Server struct {
	gin      *gin.Engine
	upgrader websocket.Upgrader
	world    *game.World
	users    map[int64]*User
}

// NewServer creates new server.
func NewServer() *Server {
	s := &Server{
		gin:      gin.Default(),
		upgrader: websocket.Upgrader{},
	}
	s.users = make(map[int64]*User)
	s.world = game.NewWorld(func(m message.Message) error {
		for _, user := range s.users {
			err := user.Write(m)
			if err != nil {
				return err
			}
		}

		return nil
	})

	s.gin.GET(
		"/ws",
		func(c *gin.Context) {
			err := s.upgrade(c.Writer, c.Request)
			if err != nil {
				log.Println(err)
				return
			}
		})

	return s
}

// Run runs server.
func (s *Server) Run() error {
	quit := make(chan bool)
	defer func() {
		quit <- true
	}()

	go func() {
		delta := time.Second / 10
		ticker := time.NewTicker(delta)

		for range ticker.C {
			select {
			case <-quit:
				return
			default:
				err := s.world.Tick(delta.Seconds())
				if err != nil {
					log.Println(err)
				}
			}
		}
	}()

	return s.gin.Run()
}

func (s *Server) upgrade(w http.ResponseWriter, r *http.Request) error {
	var upgrader websocket.Upgrader

	websocketConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}

	conn, err := net.NewConn(websocketConn)
	if err != nil {
		return err
	}

	user := NewUser(conn)

	s.users[user.ID()] = user
	defer delete(s.users, user.ID())

	err = conn.SetHandlers(net.HandlerFuncs{
		message.ActRequestID: func(conn *net.Conn, m *message.ActRequest) error {
			return ActRequestHandler(conn, m, user, s.world)
		},
		message.MoveStickRequestID: func(conn *net.Conn, m *message.MoveStickRequest) error {
			return MoveStickHandler(conn, m, user)
		},
		message.SignInRequestID: func(conn *net.Conn, m *message.SignInRequest) error {
			return SignInHandler(conn, m, user, s.world)
		},
	})
	if err != nil {
		return err
	}

	go func() {
		err := conn.Run()
		defer func() {
			user.Close()
		}()

		if err != nil {
			log.Println(err)
			return
		}
	}()

	for {
		err := conn.Consume()
		if err != nil {
			log.Println(err)
			break
		}
	}

	delete(s.users, user.ID())

	return nil
}
