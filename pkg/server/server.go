package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Server is the marsettler server.
type Server struct {
	gin *gin.Engine
}

// NewServer creates new server.
func NewServer() *Server {
	s := &Server{
		gin: gin.Default(),
	}

	s.gin.GET(
		"/ping",
		func(c *gin.Context) {
			c.JSON(
				http.StatusOK,
				gin.H{
					"message": "pong",
				})
		})

	return s
}

// Run runs server.
func (s *Server) Run() error {
	return s.gin.Run()
}
