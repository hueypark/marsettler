package heavycannon

import (
	"sync"

	"github.com/hueypark/heavycannon/body"
	"github.com/hueypark/heavycannon/contact"
	"github.com/hueypark/heavycannon/math/vector"
)

type World struct {
	bodys                 map[int64]*body.Body
	gravity               vector.Vector
	contacts              []*contact.Contact
	reservedDeleteBodyIds []int64
	mux                   sync.RWMutex
}

func New() *World {
	return &World{
		bodys:   make(map[int64]*body.Body),
		gravity: vector.Vector{X: 0.0, Y: -100.0}}
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
		b.AddForce(vector.Multiply(w.gravity, b.Mass()))
		b.Tick(delta)
	}
}

func (w *World) Add(body *body.Body) {
	w.bodys[body.Id()] = body
}

func (w *World) ReservedDelete(id int64) {
	w.reservedDeleteBodyIds = append(w.reservedDeleteBodyIds, id)
}

func (w *World) SetGravity(gravity vector.Vector) {
	w.gravity = gravity
}

func (w *World) SetBodyPosition(id int64, pos vector.Vector) {
	w.mux.Lock()
	defer w.mux.Unlock()

	b := w.bodys[id]
	if b != nil {
		b.SetPosition(pos)
	}
}

func (w *World) SetBodyVelocity(id int64, vel vector.Vector) {
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

func (w *World) Contacts() []*contact.Contact {
	return w.contacts
}

func (w *World) broadPhase() []*contact.Contact {
	contacts := []*contact.Contact{}

	for _, lhs := range w.bodys {
		for _, rhs := range w.bodys {
			if lhs.Id() <= rhs.Id() {
				continue
			}

			contacts = append(contacts, contact.New(lhs, rhs))
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
