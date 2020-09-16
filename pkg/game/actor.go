package game

import (
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/pkg/message"
)

// Actor is basic object in world.
type Actor struct {
	id                 int64
	position           vector.Vector
	moveStickDirection vector.Vector
}

// ID returns id.
func (a *Actor) ID() int64 {
	return a.id
}

// MoveStick handle move stick of actor.
func (a *Actor) MoveStick(direction vector.Vector) {
	a.moveStickDirection = direction.Normalize()
}

// Position is actor's position.
func (a *Actor) Position() vector.Vector {
	return a.position
}

// Tick updates actor periodically.
func (a *Actor) Tick(world *World, delta float64) error {
	if !a.moveStickDirection.Zero() {
		a.position = a.position.Add(a.moveStickDirection.Mul(delta))

		world.SetActorMove(&message.ActorMove{Id: a.ID(), Position: &message.Vector{X: a.position.X, Y: a.position.Y}})
	}

	return nil
}
