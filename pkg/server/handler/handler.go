package handler

import (
	"github.com/hueypark/marsettler/pkg/internal/net"
	"github.com/hueypark/marsettler/pkg/message"
	"github.com/hueypark/marsettler/pkg/server/game"
	"github.com/hueypark/marsettler/pkg/server/user"
)

// Generate generates handlers.
func Generate(user *user.User, world *game.World) net.HandlerFuncs {
	return net.HandlerFuncs{
		message.ActRequestID: func(conn *net.Conn, m *message.ActRequest) error {
			return OnActRequest(conn, m, user, world)
		},
		message.MoveStickRequestID: func(conn *net.Conn, m *message.MoveStickRequest) error {
			return OnMoveStick(conn, m, user)
		},
		message.MoveToPositionRequestID: func(conn *net.Conn, m *message.MoveToPositionRequest) error {
			return OnMoveToPositionRequest(conn, m, user)
		},
		message.SignInRequestID: func(conn *net.Conn, m *message.SignInRequest) error {
			return OnSignIn(conn, m, user, world)
		},
		message.SkillUseRequestID: func(conn *net.Conn, m *message.SkillUseRequest) error {
			return onSkillUseRequest(conn, m, user, world)
		},
	}
}
