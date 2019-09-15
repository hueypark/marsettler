package game

import (
	"github.com/hueypark/marsettler/core/id_generator"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/core/physics"
	"github.com/hueypark/marsettler/message"
)

// World represents game world.
type World struct {
	id           int64
	physicsWorld *physics.World
	actors       map[int64]*Actor
	listeners    []listener
}

// NewWorld create new world.
func NewWorld(aoe *physics.AreaOfEffect) *World {
	world := &World{
		id:           id_generator.Generate(),
		physicsWorld: physics.NewWorld(aoe),
		actors:       make(map[int64]*Actor),
	}

	return world
}

func (world *World) Actors() map[int64]*Actor {
	return world.actors
}

func (world *World) NewActor(position vector.Vector) *Actor {
	actor := NewActor(position)

	world.actors[actor.ID()] = actor
	world.physicsWorld.Add(actor)

	return actor
}

func (world *World) AddListener(l listener) {
	msgWorld := &message.World{}
	aoe := world.physicsWorld.AOE()
	msgWorld.Left = aoe.Left
	msgWorld.Right = aoe.Right
	msgWorld.Bottom = aoe.Bottom
	msgWorld.Top = aoe.Top

	l.Send(msgWorld)

	world.listeners = append(world.listeners, l)
}

// Tick ticks world.
func (world *World) Tick(delta float64) {
	msgActors := &message.Actors{}

	world.physicsWorld.Tick(delta)
	for _, actor := range world.actors {
		actor.Tick()
		msgActors.Actors = append(msgActors.Actors, &message.Actor{
			Id:  actor.ID(),
			Pos: &message.Vector{X: actor.Position().X, Y: actor.Position().Y}})
	}

	for _, l := range world.listeners {
		l.Send(msgActors)
	}
}

type listener interface {
	Send(msg message.Msg)
}
