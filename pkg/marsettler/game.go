package marsettler

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hueypark/marsettler/pkg/world"
)

// Game is the main game struct.
type Game struct {
	actors map[int]world.Actor
}

// NewGame creates a new game.
func NewGame() *Game {
	game := &Game{
		actors: make(map[int]world.Actor),
	}

	game.actors[0] = world.NewWorker()

	return game
}

// Update is the main game update loop.
func (g *Game) Update() error {
	return nil
}

// Draw is the main game draw loop.
func (g *Game) Draw(screen *ebiten.Image) {
	for _, actor := range g.actors {
		actor.Draw(screen)
	}
}

// Layout is the main game layout loop.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
