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
func NewCursor() *Cursor {
	cursor := &Cursor{}

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

// SetNode sets node.
func (cursor *Cursor) SetNode(node *game.Node) {
	cursor.node = node
}

func (cursor *Cursor) HasNode() bool {
	if cursor.node == nil {
		return false
	}

	return true
}
