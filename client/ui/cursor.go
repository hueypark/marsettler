package ui

import (
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/server/game"
)

// Cursor represent cursor.
type Cursor struct {
	nodeID int64
}

// NewCursor create new cursor.
func NewCursor(nodeID int64) *Cursor {
	cursor := &Cursor{nodeID}

	return cursor
}

// Position returns position.
func (cursor *Cursor) Position() vector.Vector {
	node := game.GetNode(cursor.nodeID)
	if node == nil {
		vector.Zero()
	}

	return node.Position()
}
