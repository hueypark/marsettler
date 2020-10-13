package handler

import (
	"errors"
	"fmt"

	"github.com/hueypark/marsettler/pkg/data"
	"github.com/hueypark/marsettler/pkg/global"
	"github.com/hueypark/marsettler/pkg/internal/math2d"
	"github.com/hueypark/marsettler/pkg/internal/net"
	"github.com/hueypark/marsettler/pkg/message"
	"github.com/hueypark/marsettler/pkg/server/game"
	"github.com/hueypark/marsettler/pkg/server/user"
)

// onSkillUseRequest handles message.SkillUseRequest.
func onSkillUseRequest(
	conn *net.Conn, m *message.SkillUseRequest, user *user.User, world *game.World,
) error {
	res := &message.SkillUseResponse{}

	actor := world.Actor(user.ID())
	if actor == nil {
		res.ResponseCode = message.ActorIsNil
		return conn.Write(res)
	}

	if m.Direction == nil {
		return errors.New(fmt.Sprintf("direction is nil [userID: %v]", user.ID()))
	}

	pos := actor.Position()

	skillActor, err := world.NewActor(global.IdGenerator.Generate().Int64(), data.SwordSkillID, pos)
	if err != nil {
		return err
	}

	dir := &math2d.Vector{X: m.Direction.X, Y: m.Direction.Y}
	dir.Normalize()
	dir.Mul(100)

	skillActor.SetMoveToPosition(math2d.Add(pos, dir))

	return conn.Write(res)
}
