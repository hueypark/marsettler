package lumberjack

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hueypark/marsettler/pkg/world"
	"github.com/unitoftime/flow/phy2"
)

type Game struct {
	world  *world.World
	lodges []any
	woods  int
}

func NewGame() *Game {
	w := world.New()
	generateDummyNodes(w)
	return &Game{
		world: w,
	}
}

func (g *Game) Run() error {
	return ebiten.RunGame(g)
}

func (g *Game) Update() error {
	return g.world.Update()
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

	g.world.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func generateDummyNodes(w *world.World) []*world.Node {
	var nodes []*world.Node
	for i := range 5 {
		for j := range 5 {
			_ = w.NewNode(phy2.V2(float64(i*100), float64(j*100)))
		}
	}

	return nodes
}
