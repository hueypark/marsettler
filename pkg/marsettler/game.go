package marsettler

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Game is the main game struct.
type Game struct {
}

// NewGame creates a new game.
func NewGame() *Game {
	return &Game{}
}

// Update is the main game update loop.
func (g *Game) Update() error {
	return nil
}

// Draw is the main game draw loop.
func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

// Layout is the main game layout loop.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
