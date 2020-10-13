package game

import (
	"fmt"

	"github.com/hueypark/marsettler/pkg/data"
	"github.com/hueypark/marsettler/pkg/internal/game"
	"github.com/hueypark/marsettler/pkg/internal/math2d"
	"github.com/hueypark/marsettler/pkg/message"
)

// Actor is basic object in world.
type Actor struct {
	*game.Actor
	moveStickDirection *math2d.Vector
	moveToPosition     *math2d.Vector
	moved              bool
}

// NewActor Creates new actor.
func NewActor(id int64, dataID data.ActorID, position *math2d.Vector) (*Actor, error) {
	a := &Actor{}

	var err error
	a.Actor, err = game.NewActor(
		id,
		dataID,
		position,
		func(position *math2d.Vector) {
			a.moved = true
		})
	if err != nil {
		return nil, err
	}

	return a, nil
}

// Act acts to target.
func (a *Actor) Act(world *World, target *Actor) error {
	return world.DeleteActor(target.ID())
}

// Message returns message.Actor.
func (a *Actor) Message() *message.Actor {
	m := &message.Actor{
		Id:       a.ID(),
		Position: &message.Vector{X: a.Position().X, Y: a.Position().Y},
		DataID:   int32(a.DataID()),
	}

	return m
}

// MoveStick handle move stick of actor.
func (a *Actor) MoveStick(direction math2d.Vector) {
	a.moveStickDirection = &direction
	a.moveStickDirection.Normalize()

	a.moveToPosition = nil
}

// SetMoveToPosition sets the position where it will move and arrive.
func (a *Actor) SetMoveToPosition(position *math2d.Vector) {
	if a.moveToPosition == nil {
		a.moveToPosition = position.Clone()
	} else {
		a.moveToPosition.Set(position)
	}

	a.moveStickDirection = nil
}

// String implements fmt.Stringer.
func (a *Actor) String() string {
	return fmt.Sprintf("id: %v, position: %v", a.ID(), a.Position())
}

// Tick updates actor periodically.
func (a *Actor) Tick(world *World, delta float64) error {
	if a.moveStickDirection != nil {
		position := a.Position()
		position.AddScaledVector(a.moveStickDirection, delta*a.Speed())

		a.SetPosition(position)

		a.moved = true
	} else if a.moveToPosition != nil {
		position := a.Position()

		moveStickDirection := math2d.Sub(a.moveToPosition, position)

		moveDistance := delta * a.Speed()
		if moveDistance*moveDistance <= moveStickDirection.SizeSquare() {
			moveStickDirection.Normalize()
			position.AddScaledVector(moveStickDirection, moveDistance)
		} else {
			position.Set(a.moveToPosition)

			a.moveToPosition = nil
		}

		a.SetPosition(position)

		a.moved = true
	}

	if a.moved {
		world.SetActorMove(&message.ActorMove{Id: a.ID(), Position: &message.Vector{X: a.Position().X, Y: a.Position().Y}})
		a.moved = false
	}

	return nil
}
