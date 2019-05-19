package ui

import (
	"github.com/hajimehoshi/ebiten"
)

// Cursor represent cursor.
type Cursor struct {
}

// NewCursor create new cursor.
func NewCursor() *Cursor {
	cursor := &Cursor{}

	return cursor
}

// Render renders cursor.
func (cursor *Cursor) Render(screen *ebiten.Image) {
	//renderer.Render(screen, asset.Cursor, cursor.Position())
}
