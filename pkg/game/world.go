package game

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	g "github.com/hueypark/marsettler/core/graph"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/data"
	"github.com/hueypark/marsettler/message"
)

// World represents game world.
type World struct {
	actors map[int64]*Actor
	graph  g.Graph
}

// NewWorld create new world.
func NewWorld() *World {
	world := &World{
		actors: make(map[int64]*Actor),
	}

	return world
}

func (w *World) NewActor(actorID data.ActorID, position, velocity vector.Vector) *Actor {
	actor := NewActor(actorID, position, velocity)

	w.actors[actor.ID()] = actor

	return actor
}

func (w *World) Actor(id int64) *Actor {
	if actor, ok := w.actors[id]; ok {
		return actor
	}

	return nil
}

func (w *World) UpsertActor(msgActor *message.Actor) {
}

func (w *World) AddListener(l listener) {
	msgWorld := &message.World{}

	l.Send(msgWorld)
}

// Tick ticks world.
func (w *World) Tick() {
}

func (w *World) Render(screen *ebiten.Image) {
	for _, actor := range w.actors {
		actor.Render(screen)
	}

	for _, iter := range w.graph.Nodes() {
		node, ok := iter.(*Node)
		if !ok {
			log.Print("Node is not game.node.")
		}

		node.Render(screen)
	}
}

type listener interface {
	Send(msg message.Msg)
}
