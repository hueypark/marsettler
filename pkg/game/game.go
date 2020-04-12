package game

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/pkg/config"
	"github.com/hueypark/marsettler/pkg/consts"
	"github.com/hueypark/marsettler/pkg/renderer"
)

// Game is a game. Entrypoint of all system.
type Game struct {
	world *World
	user  *User
}

// New create new game.
func New() *Game {
	return &Game{}
}

// Run runs game.
func (g *Game) Run() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	g.world = NewWorld()
	kingdom := NewKingdom()
	startNodeID, err := g.world.StartNodeID()
	if err != nil {
		panic(err)
	}
	g.user = g.world.NewUser(kingdom.ID(), startNodeID)

	ebiten.SetRunnableInBackground(true)
	ebiten.SetMaxTPS(consts.TPS)
	err = ebiten.Run(g.tick, config.ScreenWidth, config.ScreenHeight, 1, "Marsettler")
	if err != nil {
		log.Fatalln(err)
	}

}

func (g *Game) tick(screen *ebiten.Image) error {
	ebiten.CurrentFPS()
	x, y := ebiten.CursorPosition()
	cursorPosition := vector.Vector{X: float64(x), Y: float64(y)}

	worldPosition := renderer.WorldPosition(cursorPosition)

	g.user.Tick(worldPosition)
	g.world.Tick()
	g.tickRenderer(screen, cursorPosition)

	return ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

func (g *Game) tickRenderer(screen *ebiten.Image, cursorPosition vector.Vector) {
	_, dy := ebiten.Wheel()
	renderer.Zoom(dy*0.1, cursorPosition)

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		renderer.OnScrollStart(cursorPosition)
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		renderer.OnScrollEnd()
	}

	renderer.Tick(cursorPosition)

	g.world.Render(screen)
	g.user.Render(screen)
}
