package server

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/hueypark/marsettler/server/pkg/internal/net"
	"github.com/hueypark/marsettler/server/pkg/server/game"
	"github.com/hueypark/marsettler/server/pkg/server/handler"
	"github.com/hueypark/marsettler/server/pkg/server/user"
)

// Server is the client server.
type Server struct {
	gin		*gin.Engine
	upgrader	websocket.Upgrader
	world		*game.World
	users		[]*user.User
	mux		sync.Mutex
}

// NewServer creates new server.
func NewServer() *Server {
	s := &Server{
		gin:		gin.Default(),
		upgrader:	websocket.Upgrader{},
	}
	s.world = game.NewWorld()

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

// NewUser creates new user.
func (s *Server) NewUser(conn *net.Conn) *user.User {
	u := user.New(conn)

	s.mux.Lock()
	s.users = append(s.users, u)
	s.mux.Unlock()

	return u
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
				s.mux.Lock()
				for _, user := range s.users {
					err := user.Consume()
					if err != nil {
						log.Printf("+%v", err)
					}
				}
				s.mux.Unlock()

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

	u := s.NewUser(conn)

	err = conn.SetHandlers(
		handler.Generate(
			u,
			s.world,
		))
	if err != nil {
		return err
	}

	go func() {
		err := conn.Run()
		if err != nil {
			log.Printf("+%v", err)
			return
		}
	}()

	return nil
}
