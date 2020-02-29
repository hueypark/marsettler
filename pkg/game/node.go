package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/pkg/asset"
	"github.com/hueypark/marsettler/pkg/renderer"
)

type Node struct {
}

func (n *Node) ID() int64 {
	return 0
}

func (n *Node) Position() vector.Vector {
	return vector.Zero()
}

func (n *Node) Render(screen *ebiten.Image) {
	renderer.Render(screen, asset.Circle, n.Position())
}
