package game

import (
	"sync"

	"github.com/hueypark/marsettler/client/renderer"

	"github.com/hueypark/marsettler/core/physics"

	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/message"
)

type World struct {
	actors map[int64]*Actor
	aoes   []*physics.AreaOfEffect
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

	for _, aoe := range world.aoes {
		renderer.RenderAOE(screen, aoe)
	}
}

func (world *World) NewActor(msgActor *message.Actor) {
	world.mux.Lock()
	defer world.mux.Unlock()

	world.actors[msgActor.Id] = NewActor(msgActor)
}

func (world *World) NewAOE(aoe *physics.AreaOfEffect) {
	world.mux.Lock()
	defer world.mux.Unlock()

	world.aoes = append(world.aoes, aoe)
}
