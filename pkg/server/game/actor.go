package game

import (
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/pkg/internal/game"
	"github.com/hueypark/marsettler/pkg/message"
)

// Actor is basic object in world.
type Actor struct {
	*game.Actor
	moveStickDirection *vector.Vector
}

// NewActor Creates new actor.
func NewActor(id int64) *Actor {
	a := &Actor{
		Actor:              game.NewActor(id),
		moveStickDirection: &vector.Vector{},
	}

	return a
}

// MoveStick handle move stick of actor.
func (a *Actor) MoveStick(direction vector.Vector) {
	a.moveStickDirection = &direction
	a.moveStickDirection.Normalize()
}

// Tick updates actor periodically.
func (a *Actor) Tick(world *World, delta float64) error {
	if !a.moveStickDirection.Zero() {
		a.Position().AddScaledVector(a.moveStickDirection, delta*a.Speed())

		world.SetActorMove(&message.ActorMove{Id: a.ID(), Position: &message.Vector{X: a.Position().X, Y: a.Position().Y}})
	}

	return nil
}
