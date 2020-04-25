package game

import (
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/core/id_generator"
	"github.com/hueypark/marsettler/core/math/rotator"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/data"
	"github.com/hueypark/marsettler/pkg/asset"
	"github.com/hueypark/marsettler/pkg/consts"
	"github.com/hueypark/marsettler/pkg/renderer"
)

type Node struct {
	id        int64
	kingdomID int64
	image     *ebiten.Image
	pos       vector.Vector
	actors    map[int64]*Actor
}

// NewNode creates node.
func NewNode(pos vector.Vector) *Node {
	n := &Node{
		id:        id_generator.Generate(),
		kingdomID: 0,
		image:     asset.Image("/asset/tiles_grassland_dense_clear_green/0.png"),
		pos:       pos,
		actors:    make(map[int64]*Actor),
	}

	return n
}

// AddActor adds actor to node.
func (n *Node) AddActor(actor *Actor) {
	n.actors[actor.ID()] = actor
}

func (n *Node) ID() int64 {
	return n.id
}

// NewActor creates new actor.
func (n *Node) NewActor(kingdomID int64, actorID data.ActorID, world *World) *Actor {
	actor := newActor(kingdomID, actorID, world, n)

	n.actors[actor.ID()] = actor

	return actor
}

func (n *Node) Position() vector.Vector {
	return n.pos
}

// DeleteActor deletes actor from node.
func (n *Node) DeleteActor(id int64) {
	delete(n.actors, id)
}

func (n *Node) Render(screen *ebiten.Image) {
	pos := n.Position()
	x, _ := n.image.Size()
	radiusHalf := float64(x) * 0.5
	pos.X -= radiusHalf
	pos.Y -= radiusHalf

	kingdom, ok := Kingdoms[n.kingdomID]
	if ok {
		renderer.RenderWithColor(screen, n.image, pos, &kingdom.color)
	} else {
		renderer.Render(screen, n.image, pos)
	}
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

// SetKingdomID sets kingdom id.
func (n *Node) SetKingdomID(id int64) {
	n.kingdomID = id
}
