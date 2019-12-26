package game

import (
	"sync"

	"github.com/hajimehoshi/ebiten"
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

	mux sync.RWMutex
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

func (world *World) NewActor(id int64, position, velocity vector.Vector) *Actor {
	world.mux.Lock()
	defer world.mux.Unlock()

	actor := NewActor(id, position, velocity)

	world.actors[actor.ID()] = actor
	world.physicsWorld.Add(actor.Body())

	return actor
}

func (world *World) Actor(id int64) *Actor {
	world.mux.Lock()
	defer world.mux.Unlock()

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

	world.listeners = append(world.listeners, l)
}

// Tick ticks world.
func (world *World) Tick(delta float64) {
	world.mux.RLock()
	defer world.mux.RUnlock()

	msgActors := &message.Actors{}

	world.physicsWorld.Tick(delta)
	for _, actor := range world.actors {
		actor.Tick()
	}

	for _, l := range world.listeners {
		l.Send(msgActors)
	}
}

func (world *World) Render(screen *ebiten.Image) {
	world.mux.RLock()
	defer world.mux.RUnlock()

	for _, actor := range world.actors {
		actor.Render(screen)
	}
}

type listener interface {
	Send(msg message.Msg)
}
