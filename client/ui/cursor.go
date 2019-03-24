package ui

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/client/asset"
	"github.com/hueypark/marsettler/client/renderer"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/server/game"
)

// Cursor represent cursor.
type Cursor struct {
	node *game.Node
}

// NewCursor create new cursor.
func NewCursor(node *game.Node) *Cursor {
	cursor := &Cursor{node}

	return cursor
}

// Render renders cursor.
func (cursor *Cursor) Render(screen *ebiten.Image) {
	renderer.Render(screen, asset.Cursor, cursor.Position())
}

// Position returns position.
func (cursor *Cursor) Position() vector.Vector {
	return cursor.node.Position()
}
