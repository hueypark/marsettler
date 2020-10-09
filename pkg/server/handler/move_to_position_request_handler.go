package handler

import (
	"errors"
	"fmt"

	"github.com/hueypark/marsettler/pkg/internal/math2d"
	"github.com/hueypark/marsettler/pkg/internal/net"
	"github.com/hueypark/marsettler/pkg/message"
	"github.com/hueypark/marsettler/pkg/server/user"
)

// OnMoveToPositionRequest handles message.MoveToPosition.
func OnMoveToPositionRequest(_ *net.Conn, m *message.MoveToPositionRequest, user *user.User) error {
	actor := user.Actor()
	if actor == nil {
		return errors.New(fmt.Sprintf("actor is nil [user: %v]", user))
	}

	if m.Position == nil {
		return errors.New(fmt.Sprintf("position is nil [user: %v]", user))
	}

	actor.SetMoveToPosition(&math2d.Vector{X: m.Position.X, Y: m.Position.Y})

	return nil
}
