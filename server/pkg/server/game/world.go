package game

import (
	"errors"
	"fmt"
	"log"

	"github.com/hueypark/marsettler/server/pkg/data"
	"github.com/hueypark/marsettler/server/pkg/internal/math2d"
	"github.com/hueypark/marsettler/server/pkg/internal/physics"
	"github.com/hueypark/marsettler/server/pkg/message"
)

// World is an area where fewer than 2,000 users can play at the same time.
type World struct {
	actors				map[int64]*Actor
	physicsWorld			*physics.World
	messageActorMovesPushCache	message.ActorMovesPush
}

// NewWorld creates world.
func NewWorld() *World {
	w := &World{}
	w.actors = make(map[int64]*Actor)
	w.physicsWorld = physics.NewWorld()

	return w
}

// Actor returns actor in the world.
func (w *World) Actor(id int64) *Actor {
	return w.actors[id]
}

// ActorsPush return actors push message.
//
// It has all actor's data.
func (w *World) ActorsPush() *message.ActorsPush {
	m := &message.ActorsPush{}

	for _, actor := range w.actors {
		m.Actors = append(
			m.Actors,
			actor.Message())
	}

	return m
}

// DeleteActor deletes an actor.
func (w *World) DeleteActor(actorID int64) error {
	_, ok := w.actors[actorID]
	if !ok {
		return errors.New(fmt.Sprintf("there is no actor %v", actorID))
	}

	delete(w.actors, actorID)

	m := &message.ActorDisappearsPush{}
	m.Disappears = append(m.Disappears, &message.ActorDisappear{Id: actorID})

	return w.broadcast(m)
}

// NewActor creates new actor.
func (w *World) NewActor(id int64, dataID data.ActorID, position *math2d.Vector) (*Actor, error) {
	a, err := NewActor(id, dataID, position)
	if err != nil {
		return nil, err
	}

	_, ok := w.actors[a.ID()]
	if ok {
		return nil, errors.New(fmt.Sprintf("actor already exists [id: %v]", a.ID()))
	}

	w.actors[a.ID()] = a
	w.physicsWorld.Add(a.Body)

	m := &message.ActorsPush{}
	m.Actors = append(m.Actors, a.Message())

	err = w.broadcast(m)
	if err != nil {
		log.Println(err)
	}

	return a, nil
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

	w.physicsWorld.Tick(delta)

	w.flushActorMovePush()

	return nil
}

// broadcast sends messages to all actor.
func (w *World) broadcast(m message.Message) error {
	for _, a := range w.actors {
		err := a.Write(m)
		if err != nil {
			return err
		}
	}

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
