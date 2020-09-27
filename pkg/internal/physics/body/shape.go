package body

type Shape int

const (
	Bullet Shape = iota
	Circle
	Convex
)

type shape interface {
	Type() Shape
}
