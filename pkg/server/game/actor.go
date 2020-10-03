package game

import (
	"fmt"
	"log"

	"github.com/hueypark/marsettler/pkg/internal/game"
	"github.com/hueypark/marsettler/pkg/internal/math2d"
	"github.com/hueypark/marsettler/pkg/message"
)

// Actor is basic object in world.
type Actor struct {
	*game.Actor
	moveStickDirection *math2d.Vector
	moved              bool
}

// NewActor Creates new actor.
func NewActor(id int64) *Actor {
	a := &Actor{
		moveStickDirection: &math2d.Vector{},
	}

	a.Actor = game.NewActor(
		id,
		func(position *math2d.Vector) {
			a.moved = true
		})

	return a
}

// Act acts to target.
func (a *Actor) Act(target *Actor) error {
	log.Println(fmt.Sprintf("%v acts to %v", *a, *target))

	return nil
}

// Message returns message.Actor.
func (a *Actor) Message() *message.Actor {
	m := &message.Actor{
		Id:       a.ID(),
		Position: &message.Vector{X: a.Position().X, Y: a.Position().Y},
	}

	return m
}

// MoveStick handle move stick of actor.
func (a *Actor) MoveStick(direction math2d.Vector) {
	a.moveStickDirection = &direction
	a.moveStickDirection.Normalize()
}

// Tick updates actor periodically.
func (a *Actor) Tick(world *World, delta float64) error {
	if !a.moveStickDirection.Zero() {
		position := a.Position()
		position.AddScaledVector(a.moveStickDirection, delta*a.Speed())
		a.SetPosition(position)
		a.moved = true
	}

	if a.moved {
		world.SetActorMove(&message.ActorMove{Id: a.ID(), Position: &message.Vector{X: a.Position().X, Y: a.Position().Y}})
		a.moved = false
	}

	return nil
}
