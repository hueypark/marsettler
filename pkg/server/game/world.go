package game

import (
	"log"

	"github.com/hueypark/marsettler/pkg/message"
)

// World is an area where fewer than 2,000 users can play at the same time.
type World struct {
	actors                     map[int64]*Actor
	messageActorMovesPushCache message.ActorMovesPush
	broadcast                  func(m message.Message) error
}

// NewWorld creates world.
func NewWorld(broadcast func(m message.Message) error) *World {
	w := &World{}
	w.actors = make(map[int64]*Actor)
	w.broadcast = broadcast

	return w
}

// ActorsPush return actors push message.
//
// It has all actor's data.
func (w *World) ActorsPush() *message.ActorsPush {
	m := &message.ActorsPush{}

	for _, actor := range w.actors {
		m.Actors = append(
			m.Actors,
			&message.Actor{
				Id:       actor.ID(),
				Position: &message.Vector{X: actor.Position().X, Y: actor.Position().Y},
			})
	}

	return m
}

// NewActor creates new actor.
func (w *World) NewActor(id int64) *Actor {
	a := NewActor(id)

	w.actors[a.ID()] = a

	return a
}

// SetActorMove sets message.ActorMove message.
func (w *World) SetActorMove(m *message.ActorMove) {
	w.messageActorMovesPushCache.Moves = append(w.messageActorMovesPushCache.Moves, m)
}

// Tick updates world periodically.
func (w *World) Tick(delta float64) error {
	for _, actor := range w.actors {
		err := actor.Tick(w, delta)
		if err != nil {
			return err
		}
	}

	w.flushActorMovePush()

	return nil
}

// flushActorMovePush flushes actor move push cache to user.
func (w *World) flushActorMovePush() {
	if len(w.messageActorMovesPushCache.Moves) <= 0 {
		return
	}

	err := w.broadcast(&w.messageActorMovesPushCache)
	if err != nil {
		log.Println(err)
	}

	w.messageActorMovesPushCache.Moves = w.messageActorMovesPushCache.Moves[:0]
}
