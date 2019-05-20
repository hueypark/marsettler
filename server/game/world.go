package game

import (
	"log"

	"github.com/hueypark/marsettler/core/id_generator"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/core/physics"
	"github.com/hueypark/marsettler/data"
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

func (world *World) NewActor(id int, position vector.Vector) *Actor {
	actorData := data.Actor(id)
	if actorData == nil {
		log.Println("actor data is nil", id)
		return nil
	}

	actor := NewActor(actorData, position)

	world.actors[actor.ID()] = actor
	world.physicsWorld.AddBody(actor)

	return actor
}

// Tick ticks world.
func (world *World) Tick() {
	world.physicsWorld.Tick()
}

// RandomPath returns random path.
func (world *World) RandomPath(length int) *[]int64 {
	return nil
}
