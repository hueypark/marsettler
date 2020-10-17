package handler

import (
	"fmt"
	"log"

	"github.com/hueypark/marsettler/pkg/data"
	"github.com/hueypark/marsettler/pkg/internal/math2d"
	"github.com/hueypark/marsettler/pkg/internal/net"
	"github.com/hueypark/marsettler/pkg/message"
	"github.com/hueypark/marsettler/pkg/server/game"
	"github.com/hueypark/marsettler/pkg/server/user"
)

// OnSignIn handles message.SignIn.
func OnSignIn(conn *net.Conn, m *message.SignInRequest, user *user.User, world *game.World) error {
	response := &message.SignInResponse{}

	actor, err := world.NewActor(user.ID(), data.UserID, &math2d.Vector{})
	if err != nil {
		return err
	}

	user.SetActor(actor)

	response.Id = actor.ID()
	response.Actor = actor.Message()

	err = conn.Write(world.ActorsPush())
	if err != nil {
		return err
	}

	log.Println(fmt.Sprintf("Sign in [id: %v]", actor.ID()))

	return conn.Write(response)
}
