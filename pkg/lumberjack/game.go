package lumberjack

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	lodges []any
	woods  int
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Run() error {
	return ebiten.RunGame(g)
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf(
			`Lumberjack's lodge: %d
Woods: %d`,
			len(g.lodges),
			g.woods,
		),
		10,
		10,
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
