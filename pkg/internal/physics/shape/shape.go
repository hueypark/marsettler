package shape

type Type int

const (
	BulletType Type = iota
	CircleType
	ConvexType
)

type Shape interface {
	Type() Type
}
