package physics

import (
	"sync"

	"github.com/hueypark/marsettler/core/physics/body"
	"github.com/hueypark/marsettler/pkg/internal/math"
)

type World struct {
	bodys                 map[int64]*body.Body
	contacts              []*Contact
	reservedDeleteBodyIds []int64
	mux                   sync.RWMutex
}

func NewWorld() *World {
	return &World{
		bodys: make(map[int64]*body.Body),
	}
}

func (w *World) Tick(delta float64) {
	w.mux.Lock()
	defer w.mux.Unlock()

	w.deleteReserveDeleteBodys()

	w.contacts = w.broadPhase()
	for _, c := range w.contacts {
		c.DetectCollision()
		c.SolveCollision()
	}

	for _, b := range w.bodys {
		b.Tick(delta)
	}
}

func (w *World) Add(b *body.Body) {
	w.bodys[b.ID()] = b
}

func (w *World) ReservedDelete(id int64) {
	w.reservedDeleteBodyIds = append(w.reservedDeleteBodyIds, id)
}

func (w *World) SetBodyPosition(id int64, pos math.Vector) {
	w.mux.Lock()
	defer w.mux.Unlock()

	b := w.bodys[id]
	if b != nil {
		b.SetPosition(pos)
	}
}

func (w *World) SetBodyVelocity(id int64, vel math.Vector) {
	w.mux.Lock()
	defer w.mux.Unlock()

	b := w.bodys[id]
	if b != nil {
		b.Velocity = vel
	}
}

func (w *World) Bodys() map[int64]*body.Body {
	return w.bodys
}

func (w *World) Contacts() []*Contact {
	return w.contacts
}

func (w *World) broadPhase() []*Contact {
	var contacts []*Contact

	for _, lhs := range w.bodys {
		for _, rhs := range w.bodys {
			if lhs.ID() <= rhs.ID() {
				continue
			}

			contacts = append(contacts, New(lhs, rhs))
		}
	}

	return contacts
}

func (w *World) deleteReserveDeleteBodys() {
	for _, id := range w.reservedDeleteBodyIds {
		delete(w.bodys, id)
	}

	w.reservedDeleteBodyIds = []int64{}
}
