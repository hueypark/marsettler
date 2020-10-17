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

// ActorsPushHandler handles message.ActorsPush.
func ActorsPushHandler(conn *net.Conn, m *message.ActorsPush, world *game.World) error {
	for _, mActor := range m.Actors {
		actor, err := world.NewActor(mActor.Id, data.ActorID(mActor.DataID), &math2d.Vector{X: mActor.Position.X, Y: mActor.Position.Y})
		if err != nil {
			log.Println(err)
		}

		for _, stat := range mActor.Stats {
			if stat == nil {
				log.Println(fmt.Sprintf("stat is nil. [actorID: %v]", mActor.ID()))
				continue
			}

			switch stat.Type {
			case message.StatTypeHP:
				actor.SetHP(stat.Val)
			case message.StatTypeMaxHP:
				actor.SetMaxHP(stat.Val)
			default:
				log.Println(fmt.Sprintf("unhandled stat. [actorID: %v, type: %v]", mActor.ID(), stat.Type))
			}
		}
	}

	return nil
}
