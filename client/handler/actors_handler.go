package handler

import (
	"github.com/hueypark/marsettler/client/ctx"
	"github.com/hueypark/marsettler/message"
)

func handleActors(actors *message.Actors) error {
	for _, actor := range actors.Actors {
		ctx.World.NewActor(actor)
	}

	return nil
}
