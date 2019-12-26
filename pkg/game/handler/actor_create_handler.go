package handler

import (
	"log"

	"github.com/hueypark/marsettler/core/id_generator"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/message"
	"github.com/hueypark/marsettler/pkg/ctx"
	"github.com/hueypark/marsettler/server/config"
)

func handleActorCreate(actorCreate *message.ActorCreate) error {
	position := vector.Vector{X: actorCreate.Pos.X, Y: actorCreate.Pos.Y}
	if !ctx.World.InArea(position) {
		log.Println("not in my area ", position)
		return nil
	}

	ctx.World.NewActor(
		id_generator.Generate(),
		config.ServerID,
		position,
		vector.Zero())

	return nil
}
