package circle

import "github.com/hueypark/marsettler/pkg/internal/physics/body"

type Circle struct {
	Radius float64
}

func New(radius float64) *Circle {
	c := Circle{radius}

	return &c
}

func (c *Circle) Type() body.Shape {
	return body.Circle
}
