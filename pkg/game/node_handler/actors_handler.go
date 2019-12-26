package node_handler

import (
	"github.com/hueypark/marsettler/message"
	"github.com/hueypark/marsettler/pkg/ctx"
)

func handleActors(actors *message.Actors) error {
	for _, actor := range actors.Actors {
		ctx.World.UpsertActor(actor)
	}
	return nil
}
