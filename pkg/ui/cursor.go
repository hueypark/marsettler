package ui

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/core/math/vector"
)

// Cursor represent cursor.
type Cursor struct {
	onClick func(cursorPosition vector.Vector)
	render  func(screen *ebiten.Image, cursorPosition vector.Vector)
}

func (cursor *Cursor) Clear() {
	cursor.onClick = nil
	cursor.render = nil
}

func (cursor *Cursor) OnClick(cursorPosition vector.Vector) {
	if cursor.onClick == nil {
		return
	}

	cursor.onClick(cursorPosition)
}

// Render renders cursor.
func (cursor *Cursor) Render(screen *ebiten.Image, cursorPosition vector.Vector) {
	if cursor.render == nil {
		return
	}

	cursor.render(screen, cursorPosition)
}

func (cursor *Cursor) Set(
	onClick func(cursorPosition vector.Vector),
	render func(screen *ebiten.Image, cursorPosition vector.Vector)) {

	cursor.onClick = onClick
	cursor.render = render
}
