package handler

import (
	"github.com/hueypark/marsettler/client/ctx"
	"github.com/hueypark/marsettler/core/physics"
	"github.com/hueypark/marsettler/message"
)

func handleWorld(world *message.World) error {
	ctx.World.NewAOE(&physics.AreaOfEffect{
		Left:   world.Left,
		Right:  world.Right,
		Bottom: world.Bottom,
		Top:    world.Top,
	})
	return nil
}
