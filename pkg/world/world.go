package world

import (
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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

	body := cp.NewBody(1.0, cp.INFINITY)
	body.SetPosition(cp.Vector{X: loc.X, Y: loc.Y})
	w.space.AddBody(body)

	n := newNode(w.lastGeneratedID, loc)
	w.nodes = append(w.nodes, n)

	return n
}

func (w *World) Update() error {
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		curX, curY := ebiten.CursorPosition()
		bb := cp.NewBBForCircle(cp.Vector{X: float64(curX), Y: float64(curY)}, 10.0)
		w.space.BBQuery(
			bb,
			cp.SHAPE_FILTER_ALL,
			func(s *cp.Shape, data interface{}) {
				slog.Info("Mouse button left is just released", "position", s.Body().Position())
			},
			nil,
		)
		slog.Info("Mouse button left is just released", "Dsa", "Dsa")
	}

	return nil
}

func (w *World) Draw(screen *ebiten.Image) {
	for _, n := range w.nodes {
		n.Draw(screen)
	}
}
