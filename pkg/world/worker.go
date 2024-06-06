package world

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Worker struct {
	Location Location
}

func NewWorker() Actor {
	w := &Worker{
		Location: Location{X: 0, Y: 0},
	}

	return w
}

func (w *Worker) Draw(screen *ebiten.Image) {
	drawText(screen, "Worker", 320/2, 240/2, color.White)
}
