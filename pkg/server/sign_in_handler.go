package server

import (
	"github.com/hueypark/marsettler/pkg/game"
	"github.com/hueypark/marsettler/pkg/message"
	"github.com/hueypark/marsettler/pkg/shared"
)

// SignInHandler handles message.SignIn.
func SignInHandler(conn *shared.Conn, m *message.SignIn, user *User, world *game.World) error {
	response := &message.SignInResponse{}

	actor := world.NewActor()
	user.SetActor(actor)

	response.Id = actor.ID()
	response.Actor = &message.Actor{
		Id: actor.ID(),
		Position: &message.Vector{
			X: actor.Position().X,
			Y: actor.Position().Y,
		},
	}

	return conn.Write(response)
}
