package game

import (
	"github.com/bwmarrin/snowflake"
	"github.com/hueypark/marsettler/pkg/global"
)

// World is an area where fewer than 2,000 users can play at the same time.
type World struct {
	actors map[snowflake.ID]Actor
}

// NewWorld creates world.
func NewWorld() *World {
	w := &World{}
	w.actors = make(map[snowflake.ID]Actor)

	return w
}

// Actor creates new actor.
func (w *World) NewActor() *Actor {
	a := &Actor{
		id: global.IdGenerator.Generate().Int64(),
	}

	return a
}
