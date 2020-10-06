package physics

import (
	"log"
	"sync"
)

// World represents physics world.
type World struct {
	bodys map[int64]*Body
	mux   sync.RWMutex
}

// NewWorld creates new world.
func NewWorld() *World {
	return &World{
		bodys: make(map[int64]*Body),
	}
}

// Tick updates world periodically.
func (w *World) Tick(delta float64) {
	w.mux.Lock()
	defer w.mux.Unlock()

	contacts := w.broadPhase()
	for _, c := range contacts {
		occured, err := c.DetectCollision()
		if err != nil {
			log.Println(err)
			continue
		}

		if occured {
			c.SolveCollision()
		}
	}

	for _, b := range w.bodys {
		b.Tick(delta)
	}
}

// Add adds new body to world.
func (w *World) Add(b *Body) {
	w.bodys[b.ID()] = b
}

func (w *World) broadPhase() []*Contact {
	var contacts []*Contact

	for _, lhs := range w.bodys {
		for _, rhs := range w.bodys {
			if lhs.ID() <= rhs.ID() {
				continue
			}

			contact := New(lhs, rhs)
			if contact.lhs.Shape.Type() > contact.rhs.Shape.Type() {
				contact.swap()
			}

			contacts = append(contacts, contact)
		}
	}

	return contacts
}
