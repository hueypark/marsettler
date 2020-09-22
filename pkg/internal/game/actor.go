package game

import "github.com/hueypark/marsettler/pkg/internal/math"

// Actor is basic object in world.
type Actor struct {
	id       int64
	position *math.Vector
	speed    float64
}

// NewActor creates new actor.
func NewActor(id int64) *Actor {
	a := &Actor{
		id:       id,
		position: &math.Vector{},
		speed:    100,
	}

	return a
}

// ID returns id.
func (a *Actor) ID() int64 {
	return a.id
}

// Position is actor's position.
func (a *Actor) Position() *math.Vector {
	return a.position
}

// SetPosition sets position.
func (a *Actor) SetPosition(position *math.Vector) {
	a.position = position
}

// Speed returns speed of actor.
func (a *Actor) Speed() float64 {
	return a.speed
}
