package handler

import (
	"fmt"
	"log"

	"github.com/hueypark/marsettler/pkg/client/game"
	"github.com/hueypark/marsettler/pkg/data"
	"github.com/hueypark/marsettler/pkg/internal/math2d"
	"github.com/hueypark/marsettler/pkg/internal/net"
	"github.com/hueypark/marsettler/pkg/message"
)

// SignInResponseHandler handles message.SignInResponse.
func SignInResponseHandler(
	_ *net.Conn, m *message.SignInResponse, c client, world *game.World,
) error {
	actor, err := world.NewActor(m.Actor.Id, data.UserID, &math2d.Vector{X: m.Actor.Position.X, Y: m.Actor.Position.Y})
	if err != nil {
		return err
	}

	for _, stat := range m.Actor.Stats {
		if stat == nil {
			log.Println(fmt.Sprintf("stat is nil. [actorID: %v]", m.Actor.ID()))
			continue
		}

		switch stat.Type {
		case message.StatTypeHP:
			actor.SetHP(stat.Val)
		case message.StatTypeMaxHP:
			actor.SetMaxHP(stat.Val)
		default:
			log.Println(fmt.Sprintf("unhandled stat. [actorID: %v, type: %v]", m.Actor.ID(), stat.Type))
		}
	}

	c.SetMyActor(actor)

	return nil
}
