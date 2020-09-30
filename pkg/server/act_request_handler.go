package server

import (
	"github.com/hueypark/marsettler/pkg/internal/net"
	"github.com/hueypark/marsettler/pkg/message"
	"github.com/hueypark/marsettler/pkg/server/game"
)

// ActRequestHandler handles message.ActRequest.
func ActRequestHandler(conn *net.Conn, m *message.ActRequest, user *User, world *game.World) error {
	response := &message.ActResponse{}

	actor := user.Actor()
	if actor == nil {
		response.ResponseCode = message.ResponseCode_ActorIsNil
		return conn.Write(response)
	}

	target := world.Actor(m.TargetId)
	if target == nil {
		response.ResponseCode = message.ResponseCode_TargetIsNil
		return conn.Write(response)
	}

	err := actor.Act(target)
	if err != nil {
		return err
	}

	return conn.Write(response)
}
