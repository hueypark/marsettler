package handler

import (
	"github.com/hueypark/marsettler/server/pkg/internal/net"
	"github.com/hueypark/marsettler/server/pkg/message"
	"github.com/hueypark/marsettler/server/pkg/server/game"
	"github.com/hueypark/marsettler/server/pkg/server/user"
)

// OnActRequest handles message.ActRequest.
func OnActRequest(conn *net.Conn, m *message.ActRequest, user *user.User, world *game.World) error {
	response := &message.ActResponse{}

	actor := user.Actor()
	if actor == nil {
		response.ResponseCode = message.ActorIsNil
		return conn.Write(response)
	}

	target := world.Actor(m.TargetId)
	if target == nil {
		response.ResponseCode = message.TargetIsNil
		return conn.Write(response)
	}

	err := actor.Act(world, target)
	if err != nil {
		return err
	}

	return conn.Write(response)
}
