package game

import (
	"github.com/hueypark/marsettler/core/id_generator"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/core/physics"
)

// World represents game world.
type World struct {
	id           int64
	physicsWorld *physics.World
	actors       map[int64]*Actor
}

// NewWorld create new world.
func NewWorld() *World {
	world := &World{
		id:           id_generator.Generate(),
		physicsWorld: physics.NewWorld(),
		actors:       make(map[int64]*Actor),
	}

	return world
}

func (world *World) Actors() map[int64]*Actor {
	return world.actors
}

func (world *World) NewActor(id int64, position vector.Vector) *Actor {
	actor := NewActor(position)

	world.actors[actor.ID()] = actor
	world.physicsWorld.AddBody(actor)

	return actor
}

// Tick ticks world.
func (world *World) Tick() {
	// TODO: Send actor information to all users in this world.

	world.physicsWorld.Tick()
	for _, actor := range world.actors {
		actor.Tick()
	}
}
