package handler

import (
	"github.com/hueypark/marsettler/pkg/client/game"
	"github.com/hueypark/marsettler/pkg/internal/net"
	"github.com/hueypark/marsettler/pkg/message"
)

// Generate generates handlers.
func Generate(c client, world *game.World) net.HandlerFuncs {
	return net.HandlerFuncs{
		message.ActorDisappearsPushID: func(conn *net.Conn, m *message.ActorDisappearsPush) error {
			return ActorDisappearsPushHandler(conn, m, world)
		},
		message.ActResponseID: func(conn *net.Conn, m *message.ActResponse) error {
			return ActResponseHandler(conn, m)
		},
		message.ActorMovesPushID: func(conn *net.Conn, m *message.ActorMovesPush) error {
			return ActorMovesPushHandler(conn, m, world)
		},
		message.ActorsPushID: func(conn *net.Conn, m *message.ActorsPush) error {
			return ActorsPushHandler(conn, m, world)
		},
		message.SignInResponseID: func(conn *net.Conn, m *message.SignInResponse) error {
			return SignInResponseHandler(conn, m, c, world)
		},
	}
}
