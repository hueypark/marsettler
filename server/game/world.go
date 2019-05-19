package game

import (
	"github.com/hueypark/marsettler/core/id_generator"
	"github.com/hueypark/marsettler/core/physics"
)

// World represents game world.
type World struct {
	id           int64
	physicsWorld *physics.World
}

// NewWorld create new world.
func NewWorld() *World {
	world := &World{
		id:           id_generator.Generate(),
		physicsWorld: physics.NewWorld(),
	}

	return world
}

// Tick ticks world.
func (world *World) Tick() {
}

// RandomPath returns random path.
func (world *World) RandomPath(length int) *[]int64 {
	return nil
}
