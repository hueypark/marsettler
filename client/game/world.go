package game

import (
	"log"

	"github.com/hueypark/marsettler/message"
)

type World struct {
	actors map[int64]*Actor
}

// NewWorld creates new world.
func NewWorld() *World {
	world := &World{
		actors: make(map[int64]*Actor),
	}

	return world
}

func (world *World) NewActor(msgActor *message.Actor) {
	log.Fatalln(msgActor)
}
