package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/core/math/vector"
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

// NewActor creates new actor.
func (w *World) NewActor(id int64, position vector.Vector) *Actor {
	a := NewActor(id)

	w.actors[a.ID()] = a

	return a
}

func (w *World) Actor(id int64) *Actor {
	return w.actors[id]
}

// Draw drasw world.
func (w *World) Draw(screen *ebiten.Image) {
	for _, a := range w.actors {
		a.Draw(screen)
	}
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
