package client

import (
	"github.com/hueypark/marsettler/pkg/client/game"
	"github.com/hueypark/marsettler/pkg/internal/math2d"
	"github.com/hueypark/marsettler/pkg/internal/net"
	"github.com/hueypark/marsettler/pkg/message"
)

// ActorMovesPushHandler handles message.ActorMovesPush.
func ActorMovesPushHandler(conn *net.Conn, m *message.ActorMovesPush, world *game.World) error {
	for _, move := range m.Moves {
		actor := world.Actor(move.Id)
		if actor == nil {
			continue
		}

		actor.SetPosition(&math2d.Vector{X: move.Position.X, Y: move.Position.Y})
	}

	return nil
}
