package ui

import (
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

// Position returns position.
func (cursor *Cursor) Position() vector.Vector {
	return cursor.node.Position()
}
