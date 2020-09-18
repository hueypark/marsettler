package server

import (
	"errors"
	"fmt"
	"log"

	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/pkg/message"
	"github.com/hueypark/marsettler/pkg/shared"
)

// MoveStickHandler handles message.ActorMove.
func MoveStickHandler(conn *shared.Conn, m *message.MoveStickRequest, user *User) error {
	actor := user.Actor()
	if actor == nil {
		log.Println(fmt.Sprintf("actor is nil. [user: %v]", user.ID()))
		return nil
	}

	if m.Direction == nil {
		return errors.New("direction is nil")
	}

	actor.MoveStick(vector.Vector{X: m.Direction.X, Y: m.Direction.Y})

	return nil
}
