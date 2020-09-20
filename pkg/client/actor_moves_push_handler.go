package client

import (
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/pkg/client/game"
	"github.com/hueypark/marsettler/pkg/message"
	"github.com/hueypark/marsettler/pkg/shared"
)

// ActorMovesPushHandler handles message.ActorMovesPush.
func ActorMovesPushHandler(conn *shared.Conn, m *message.ActorMovesPush, world *game.World) error {
	for _, move := range m.Moves {
		actor := world.Actor(move.Id)
		if actor == nil {
			continue
		}

		actor.SetPosition(&vector.Vector{X: move.Position.X, Y: move.Position.Y})
	}

	return nil
}
