package handler

import (
	"github.com/hueypark/marsettler/server/pkg/internal/net"
	"github.com/hueypark/marsettler/server/pkg/message"
	"github.com/hueypark/marsettler/server/pkg/message/fbs"
	"github.com/hueypark/marsettler/server/pkg/server/game"
	"github.com/hueypark/marsettler/server/pkg/server/user"
)

// Generate generates handlers.
func Generate(user *user.User, world *game.World) net.HandlerFuncs {
	return net.HandlerFuncs{
		fbs.SignInRequestID: func(conn *net.Conn, m *message.SignInRequest) error {
			return OnSignIn(conn, m, user, world)
		},
	}
}
