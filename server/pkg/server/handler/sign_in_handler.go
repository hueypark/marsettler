package handler

import (
	"fmt"
	"log"

	"github.com/hueypark/marsettler/server/pkg/data"
	"github.com/hueypark/marsettler/server/pkg/global"
	"github.com/hueypark/marsettler/server/pkg/internal/math2d"
	"github.com/hueypark/marsettler/server/pkg/internal/net"
	"github.com/hueypark/marsettler/server/pkg/message"
	"github.com/hueypark/marsettler/server/pkg/message/fbs"
	"github.com/hueypark/marsettler/server/pkg/server/game"
	"github.com/hueypark/marsettler/server/pkg/server/user"
)

// OnSignIn handles message.SignIn.
func OnSignIn(conn *net.Conn, m *fbs.SignInRequest, u *user.User, world *game.World) error {
	response := &message.SignInResponse{}

	var userID int64
	if m.Id() == 0 {
		userID = global.IdGenerator.Generate().Int64()
	} else {
		userID = m.Id()
	}

	u.SetID(userID)

	actor := world.Actor(userID)
	if actor == nil {
		var err error
		actor, err = world.NewActor(userID, data.UserID, &math2d.Vector{})
		if err != nil {
			return err
		}
	}

	u.SetActor(actor)

	response.Id = actor.ID()
	response.Actor = actor.Message()

	err := conn.Write(world.ActorsPush())
	if err != nil {
		return err
	}

	log.Println(fmt.Sprintf("Sign in [id: %v]", userID))

	return conn.Write(response)
}
