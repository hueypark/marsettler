package game

import "github.com/hueypark/marsettler/core/math/vector"

// Actor is basic object in world.
type Actor struct {
	id       int64
	position *vector.Vector
	speed    float64
}

// NewActor creates new actor.
func NewActor(id int64) *Actor {
	a := &Actor{
		id:       id,
		position: &vector.Vector{},
		speed:    100,
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

// SetPosition sets position.
func (a *Actor) SetPosition(position *vector.Vector) {
	a.position = position
}

// Speed returns speed of actor.
func (a *Actor) Speed() float64 {
	return a.speed
}
