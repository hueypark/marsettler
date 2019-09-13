package game

import (
	"sync"

	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/message"
)

type World struct {
	actors map[int64]*Actor
	mux    sync.Mutex
}

// NewWorld creates new world.
func NewWorld() *World {
	world := &World{
		actors: make(map[int64]*Actor),
	}

	return world
}

func (world *World) Tick(screen *ebiten.Image) {
	for _, actor := range world.actors {
		actor.Render(screen)
	}
}

func (world *World) NewActor(msgActor *message.Actor) {
	world.mux.Lock()
	defer world.mux.Unlock()

	world.actors[msgActor.Id] = NewActor(msgActor)
}
