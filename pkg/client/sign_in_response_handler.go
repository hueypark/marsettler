package client

import (
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/pkg/client/game"
	"github.com/hueypark/marsettler/pkg/internal/net"
	"github.com/hueypark/marsettler/pkg/message"
)

// SignInResponseHandler handles message.SignInResponse.
func SignInResponseHandler(conn *net.Conn, m *message.SignInResponse, world *game.World) error {
	_ = world.NewActor(m.Actor.Id, vector.Vector{X: m.Actor.Position.X, Y: m.Actor.Position.Y})

	return nil
}
