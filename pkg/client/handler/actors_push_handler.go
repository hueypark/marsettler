package handler

import (
	"log"

	"github.com/hueypark/marsettler/pkg/client/game"
	"github.com/hueypark/marsettler/pkg/data"
	"github.com/hueypark/marsettler/pkg/internal/math2d"
	"github.com/hueypark/marsettler/pkg/internal/net"
	"github.com/hueypark/marsettler/pkg/message"
)

// ActorsPushHandler handles message.ActorsPush.
func ActorsPushHandler(conn *net.Conn, m *message.ActorsPush, world *game.World) error {
	for _, actor := range m.Actors {
		_, err := world.NewActor(actor.Id, data.ActorID(actor.DataID), &math2d.Vector{X: actor.Position.X, Y: actor.Position.Y})
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}
