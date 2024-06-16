package world

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/unitoftime/flow/phy2"
)

type Node struct {
	ID       int
	Location phy2.Vec
}

func newNode(id int, loc phy2.Vec) *Node {
	return &Node{
		ID:       id,
		Location: loc,
	}
}

var nodeImage = ebiten.NewImage(50, 50)

func init() {
	nodeImage.Fill(color.RGBA{
		R: 100,
		G: 100,
		B: 100,
		A: 100,
	})
}

func (n *Node) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(n.Location.X, n.Location.Y)
	screen.DrawImage(nodeImage, op)

	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf("Node %d", n.ID),
		int(n.Location.X),
		int(n.Location.Y),
	)
}
