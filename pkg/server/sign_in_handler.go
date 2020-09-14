package server

import (
	"github.com/hueypark/marsettler/pkg/message"
	"github.com/hueypark/marsettler/pkg/shared"
)

// SignInHandler handles message.SignIn.
func SignInHandler(conn *shared.Conn, m *message.SignIn) error {
	response := &message.SignInResponse{}
	response.Id = IdGenerator.Generate().Int64()
	response.Position = &message.Vector{
		X: 0,
		Y: 0,
	}

	return conn.Write(response)
}
