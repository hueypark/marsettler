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
	"github.com/hueypark/marsettler/pkg/ui"
)

// Game is a game. Entrypoint of all system.
type Game struct {
	world *World
	user  *User
	ui    ui.Layer
}

// New create new game.
func New() *Game {
	ebiten.SetRunnableInBackground(true)
	ebiten.SetMaxTPS(consts.TPS)

	world := NewWorld()
	kingdom := NewKingdom()
	startNodeID, err := world.StartNodeID()
	if err != nil {
		panic(err)
	}

	game := &Game{
		world: world,
		user:  world.NewUser(kingdom.ID(), startNodeID),
		ui: ui.NewButton(
			"/asset/ui/button.png",
			vector.Vector{X: 120, Y: float64(config.ScreenHeight - 70)},
			func() {
				log.Println("Button clicked.")
			}),
	}

	return game
}

// Layout implements the ebiten.Game interface.
func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return config.ScreenWidth, config.ScreenHeight
}

// Update implements the ebiten.Game interface.
func (g *Game) Update(screen *ebiten.Image) error {
	ebiten.CurrentFPS()
	x, y := ebiten.CursorPosition()
	cursorPosition := vector.Vector{X: float64(x), Y: float64(y)}

	worldPosition := renderer.WorldPosition(cursorPosition)

	g.user.Tick(worldPosition)
	g.world.Tick()
	g.tickRenderer(screen, cursorPosition)

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.onClick(cursorPosition)
	}

	return ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

func (g *Game) onClick(cursorPosition vector.Vector) {
	if g.ui != nil {
		g.ui.OnClick(cursorPosition)
	}
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

	if g.ui != nil {
		g.ui.Render(screen, nil)
	}
}
