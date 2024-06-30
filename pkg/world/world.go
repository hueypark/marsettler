package world

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp/v2"
	"github.com/unitoftime/flow/phy2"
)

type World struct {
	space           *cp.Space
	lastGeneratedID int
	nodes           []*Node
}

func New() *World {
	return &World{
		space:           cp.NewSpace(),
		lastGeneratedID: -1,
	}
}

func (w *World) NewNode(loc phy2.Vec) *Node {
	w.lastGeneratedID++
	n := newNode(w.lastGeneratedID, loc)
	w.nodes = append(w.nodes, n)

	return n
}

func (w *World) Draw(screen *ebiten.Image) {
	for _, n := range w.nodes {
		n.Draw(screen)
	}
}
