package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/pkg/asset"
	"github.com/hueypark/marsettler/pkg/renderer"
)

type Node struct {
	image *ebiten.Image
}

// NewNode creates node.
func NewNode() *Node {
	n := &Node{}
	n.image = asset.Image("grassland")

	return n
}

func (n *Node) ID() int64 {
	return 0
}

func (n *Node) Position() vector.Vector {
	return vector.Zero()
}

func (n *Node) Render(screen *ebiten.Image) {
	renderer.Render(screen, n.image, n.Position())
}
