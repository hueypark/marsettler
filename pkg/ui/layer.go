package ui

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/core/math/vector"
)

// Layer is the base object of the user interface.
type Layer interface {
	OnClick(cursorPosition vector.Vector) bool
	Render(screen *ebiten.Image, relativePos *vector.Vector)
}
