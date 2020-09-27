package shape

type Type int

const (
	BulletType Type = iota
	CircleType
	ConvecType
)

type Shape interface {
	Type() Type
}
