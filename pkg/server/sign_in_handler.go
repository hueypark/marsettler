package server

import (
	"fmt"
	"log"

	"github.com/hueypark/marsettler/pkg/internal/net"
	"github.com/hueypark/marsettler/pkg/message"
	"github.com/hueypark/marsettler/pkg/server/game"
)

// SignInHandler handles message.SignIn.
func SignInHandler(conn *net.Conn, m *message.SignInRequest, user *User, world *game.World) error {
	response := &message.SignInResponse{}

	actor, err := world.NewActor(user.ID())
	if err != nil {
		return err
	}

	user.SetActor(actor)

	response.Id = actor.ID()
	response.Actor = &message.Actor{
		Id: actor.ID(),
		Position: &message.Vector{
			X: actor.Position().X,
			Y: actor.Position().Y,
		},
	}

	err = conn.Write(world.ActorsPush())
	if err != nil {
		return err
	}

	log.Println(fmt.Sprintf("Sign in [id: %v]", actor.ID()))

	return conn.Write(response)
}
