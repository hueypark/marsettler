package game

import "github.com/hueypark/marsettler/core/math/vector"

// Actor is basic object in world.
type Actor struct {
	id       int64
	position *vector.Vector
}

// NewActor creates new actor.
func NewActor(id int64) *Actor {
	a := &Actor{
		id:       id,
		position: &vector.Vector{},
	}

	return a
}

// ID returns id.
func (a *Actor) ID() int64 {
	return a.id
}

// Position is actor's position.
func (a *Actor) Position() *vector.Vector {
	return a.position
}
