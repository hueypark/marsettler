package game

import (
	"log"
	"sync"

	"github.com/hueypark/marsettler/core/id_generator"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/core/physics"
	"github.com/hueypark/marsettler/message"
	"github.com/hueypark/marsettler/server/config"
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
func NewWorld(aoe *physics.AreaOfEffect) *World {
	world := &World{
		id:           id_generator.Generate(),
		physicsWorld: physics.NewWorld(aoe),
		actors:       make(map[int64]*Actor),
	}

	return world
}

func (world *World) NewActor(id int64, serverID int32, position, velocity vector.Vector) *Actor {
	world.mux.Lock()
	defer world.mux.Unlock()

	actor := NewActor(id, serverID, position, velocity)

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
	if actor := world.Actor(msgActor.Id); actor != nil {
		if actor.ServerID() == config.ServerID {
			return
		}

		actor.SetServerID(msgActor.ServerId)
		actor.SetPosition(vector.Vector{X: msgActor.Position.X, Y: msgActor.Position.Y})
		actor.SetVelocity(vector.Vector{X: msgActor.Velocity.X, Y: msgActor.Velocity.Y})
	} else {
		world.NewActor(
			msgActor.Id,
			msgActor.ServerId,
			vector.Vector{X: msgActor.Position.X, Y: msgActor.Position.Y},
			vector.Vector{X: msgActor.Velocity.X, Y: msgActor.Velocity.Y})
	}
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
	world.mux.RLock()
	defer world.mux.RUnlock()

	msgActors := &message.Actors{}

	world.physicsWorld.Tick(delta)
	for _, actor := range world.actors {
		actor.Tick()

		if world.InArea(actor.Position()) {
			actor.SetServerID(config.ServerID)
		} else {
			actor.SetServerID(-1)
		}

		if actor.ServerID() == 0 {
			log.Println("I'm 0")
		}

		msgActors.Actors = append(msgActors.Actors, &message.Actor{
			Id:       actor.ID(),
			ServerId: actor.ServerID(),
			Position: &message.Vector{X: actor.Position().X, Y: actor.Position().Y},
			Velocity: &message.Vector{X: actor.Velocity().X, Y: actor.Velocity().Y}})
	}

	for _, l := range world.listeners {
		l.Send(msgActors)
	}
}

func (world *World) InArea(position vector.Vector) bool {
	return world.physicsWorld.AOE().InArea(position)
}

type listener interface {
	Send(msg message.Msg)
}
