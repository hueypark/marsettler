package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/data"
	"github.com/hueypark/marsettler/message"
)

// World represents game world.
type World struct {
	actors map[int64]*Actor
}

// NewWorld create new world.
func NewWorld() *World {
	world := &World{
		actors: make(map[int64]*Actor),
	}

	return world
}

func (world *World) NewActor(actorID data.ActorID, position, velocity vector.Vector) *Actor {
	actor := NewActor(actorID, position, velocity)

	world.actors[actor.ID()] = actor

	return actor
}

func (world *World) Actor(id int64) *Actor {
	if actor, ok := world.actors[id]; ok {
		return actor
	}

	return nil
}

func (world *World) UpsertActor(msgActor *message.Actor) {
}

func (world *World) AddListener(l listener) {
	msgWorld := &message.World{}

	l.Send(msgWorld)
}

// Tick ticks world.
func (world *World) Tick() {
}

func (world *World) Render(screen *ebiten.Image) {
	for _, actor := range world.actors {
		actor.Render(screen)
	}
}

type listener interface {
	Send(msg message.Msg)
}
