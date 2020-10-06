package shape

type Type int

const (
	BulletType Type = iota
	CircleType
	ConvexType
	WallType
)

type Shape interface {
	Type() Type
}
