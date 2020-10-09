package game

import (
	"github.com/hueypark/marsettler/pkg/internal/math2d"
	"github.com/hueypark/marsettler/pkg/internal/physics"
	"github.com/hueypark/marsettler/pkg/internal/physics/shape"
)

// Actor is basic object in world.
type Actor struct {
	*physics.Body
	speed float64
}

// NewActor creates new actor.
func NewActor(
	id int64, position *math2d.Vector, onSetPosition func(position *math2d.Vector),
) *Actor {
	a := &Actor{
		Body:  physics.NewBody(id, position, onSetPosition),
		speed: 100,
	}

	a.Body.SetShape(shape.NewCircle(16))
	a.Body.SetMass(10)

	return a
}

// Speed returns speed of actor.
func (a *Actor) Speed() float64 {
	return a.speed
}
