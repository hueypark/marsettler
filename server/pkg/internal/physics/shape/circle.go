package shape

// Circle represents physics shape circle.
type Circle struct {
	Radius float64
}

// NewCircle creates new circle.
func NewCircle(radius float64) *Circle {
	c := Circle{radius}

	return &c
}

// Type returns type.
func (c *Circle) Type() Type {
	return CircleType
}
