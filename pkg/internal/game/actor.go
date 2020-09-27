package game

import (
	"github.com/hueypark/marsettler/pkg/internal/math2d"
	"github.com/hueypark/marsettler/pkg/internal/physics/body"
	"github.com/hueypark/marsettler/pkg/internal/physics/body/circle"
)

// Actor is basic object in world.
type Actor struct {
	*body.Body
	speed float64
}

// NewActor creates new actor.
func NewActor(id int64) *Actor {
	a := &Actor{
		Body:  body.New(id, &math2d.Vector{}),
		speed: 100,
	}

	a.Body.SetShape(circle.New(100))
	a.Body.SetMass(10)

	return a
}

// Speed returns speed of actor.
func (a *Actor) Speed() float64 {
	return a.speed
}
