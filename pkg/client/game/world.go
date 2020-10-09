package game

import (
	"errors"
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/pkg/internal/math2d"
)

// World represents game world.
type World struct {
	actors map[int64]*Actor
}

// NewWorld creates new world.
func NewWorld() *World {
	w := &World{}
	w.actors = make(map[int64]*Actor)

	return w
}

// DeleteActor deletes an actor.
func (w *World) DeleteActor(id int64) error {
	_, ok := w.actors[id]
	if !ok {
		return errors.New(fmt.Sprintf("there is no actor %v", id))
	}

	delete(w.actors, id)

	return nil
}

// NewActor creates new actor.
func (w *World) NewActor(id int64, position *math2d.Vector) *Actor {
	a := NewActor(id, position)

	w.actors[a.ID()] = a

	return a
}

func (w *World) Actor(id int64) *Actor {
	return w.actors[id]
}

// Draw drasw world.
func (w *World) Draw(screen *ebiten.Image, cameraFunc func(*Actor) ebiten.GeoM) error {
	for _, a := range w.actors {
		err := a.Draw(screen, cameraFunc)
		if err != nil {
			return err
		}
	}

	return nil
}

// NearestActor returns the nearest actor from the given actor.
func (w *World) NearestActor(givenActorID int64) (actor *Actor, err error) {
	givenActor := w.Actor(givenActorID)
	if givenActor == nil {
		return nil, errors.New(fmt.Sprintf("actor is nil. [givenActorId: %v]", givenActorID))
	}

	var nearestActor *Actor
	nearestDistanceSquared := math.MaxFloat64
	for _, actor := range w.actors {
		if actor.ID() == givenActorID {
			continue
		}

		distanceSquared := math2d.DistanceSquared(givenActor.Position(), actor.Position())
		if distanceSquared <= nearestDistanceSquared {
			nearestActor = actor
			nearestDistanceSquared = distanceSquared
		}
	}

	if nearestActor == nil {
		return nil, errors.New("there is no nearest actor")
	}

	return nearestActor, nil
}

// Tick updates world periodically.
func (w *World) Tick(delta float64) error {
	for _, actor := range w.actors {
		err := actor.Tick(delta)
		if err != nil {
			return err
		}
	}

	return nil
}
