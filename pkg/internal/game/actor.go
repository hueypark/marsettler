package game

import (
	"errors"
	"fmt"

	"github.com/hueypark/marsettler/pkg/data"
	"github.com/hueypark/marsettler/pkg/internal/math2d"
	"github.com/hueypark/marsettler/pkg/internal/physics"
	"github.com/hueypark/marsettler/pkg/internal/physics/shape"
)

// Actor is basic object in world.
type Actor struct {
	*physics.Body
	dataID data.ActorID
	speed  float64
}

// NewActor creates new actor.
func NewActor(
	id int64,
	dataID data.ActorID,
	position *math2d.Vector,
	onSetPosition func(position *math2d.Vector),
) (*Actor, error) {
	data := data.Actor(dataID)
	if data == nil {
		return nil, errors.New(fmt.Sprintf("no data [dataID: %v]", dataID))
	}

	a := &Actor{
		Body:   physics.NewBody(id, position, onSetPosition),
		dataID: dataID,
		speed:  100,
	}

	a.Body.SetShape(shape.NewCircle(data.Radius))
	a.Body.SetMass(10)

	return a, nil
}

// DataID returns data id.
func (a *Actor) DataID() data.ActorID {
	return a.dataID
}

// Speed returns speed of actor.
func (a *Actor) Speed() float64 {
	return a.speed
}
