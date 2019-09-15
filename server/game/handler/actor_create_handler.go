package handler

import (
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/message"
	"github.com/hueypark/marsettler/server/ctx"
)

func handleActorCreate(actorCreate *message.ActorCreate) error {
	ctx.World.NewActor(vector.Vector{X: actorCreate.Pos.X, Y: actorCreate.Pos.Y})

	return nil
}
