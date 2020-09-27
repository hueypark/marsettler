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
func NewActor(id int64) *Actor {
	a := &Actor{
		Body:  physics.NewBody(id, &math2d.Vector{}),
		speed: 100,
	}

	a.Body.SetShape(shape.NewCircle(100))
	a.Body.SetMass(10)

	return a
}

// Speed returns speed of actor.
func (a *Actor) Speed() float64 {
	return a.speed
}
