package shape

import "github.com/hueypark/marsettler/pkg/internal/math2d"

// Wall represents physics shape wall.
type Wall struct {
	start *math2d.Vector
	end   *math2d.Vector
}

// NewWall creates new wall.
func NewWall(start, end *math2d.Vector) *Wall {
	return &Wall{start, end}
}

// Type returns type.
func (w *Wall) Type() Type {
	return WallType
}
