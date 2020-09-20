package client

import (
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/pkg/client/game"
	"github.com/hueypark/marsettler/pkg/message"
	"github.com/hueypark/marsettler/pkg/shared"
)

// SignInResponseHandler handles message.SignInResponse.
func SignInResponseHandler(conn *shared.Conn, m *message.SignInResponse, world *game.World) error {
	_ = world.NewActor(m.Actor.Id, vector.Vector{X: m.Actor.Position.X, Y: m.Actor.Position.Y})

	return nil
}
