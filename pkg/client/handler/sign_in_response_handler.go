package handler

import (
	"github.com/hueypark/marsettler/pkg/client/game"
	"github.com/hueypark/marsettler/pkg/internal/math2d"
	"github.com/hueypark/marsettler/pkg/internal/net"
	"github.com/hueypark/marsettler/pkg/message"
)

// SignInResponseHandler handles message.SignInResponse.
func SignInResponseHandler(
	_ *net.Conn, m *message.SignInResponse, c client, world *game.World,
) error {
	actor := world.NewActor(m.Actor.Id, &math2d.Vector{X: m.Actor.Position.X, Y: m.Actor.Position.Y})
	c.SetMyActor(actor)

	return nil
}
