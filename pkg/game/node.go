package game

import (
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/core/asset"
	"github.com/hueypark/marsettler/core/id_generator"
	"github.com/hueypark/marsettler/core/math/rotator"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/pkg/consts"
	"github.com/hueypark/marsettler/pkg/renderer"
)

type Node struct {
	id    int64
	image *ebiten.Image
	pos   vector.Vector
}

// NewNode creates node.
func NewNode(pos vector.Vector) *Node {
	n := &Node{
		id:    id_generator.Generate(),
		image: asset.Image("/asset/tiles_grassland_dense_clear_green/0.png"),
		pos:   pos,
	}

	return n
}

func (n *Node) ID() int64 {
	return n.id
}

func (n *Node) Position() vector.Vector {
	return n.pos
}

func (n *Node) Render(screen *ebiten.Image) {
	renderer.Render(screen, n.image, n.pos)
}

func (n *Node) GetNeighborNodePositions() [6]vector.Vector {
	vec := vector.Vector{X: consts.NodeSize, Y: 0}
	rot := rotator.NewRotator(math.Pi / 3.0)

	var positions [6]vector.Vector
	positions[0] = n.pos.Add(vec)
	vec = rot.RotateVector(vec)
	positions[1] = n.pos.Add(vec)
	vec = rot.RotateVector(vec)
	positions[2] = n.pos.Add(vec)
	vec = rot.RotateVector(vec)
	positions[3] = n.pos.Add(vec)
	vec = rot.RotateVector(vec)
	positions[4] = n.pos.Add(vec)
	vec = rot.RotateVector(vec)
	positions[5] = n.pos.Add(vec)

	return positions
}
