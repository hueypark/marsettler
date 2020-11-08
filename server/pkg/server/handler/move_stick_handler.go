// +build ignore

package handler

import (
	"errors"
	"fmt"
	"log"

	"github.com/hueypark/marsettler/server/pkg/internal/math2d"
	"github.com/hueypark/marsettler/server/pkg/internal/net"
	"github.com/hueypark/marsettler/server/pkg/message"
	"github.com/hueypark/marsettler/server/pkg/server/user"
)

// OnMoveStick handles message.ActorMove.
func OnMoveStick(conn *net.Conn, m *message.MoveStickRequest, user *user.User) error {
	actor := user.Actor()
	if actor == nil {
		log.Println(fmt.Sprintf("actor is nil. [user: %v]", user.ID()))
		return nil
	}

	if m.Direction == nil {
		return errors.New("direction is nil")
	}

	actor.MoveStick(math2d.Vector{X: m.Direction.X, Y: m.Direction.Y})

	return nil
}
