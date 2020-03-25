package bullet

import "github.com/hueypark/marsettler/core/physics/body"

type Bullet struct {
}

func New() *Bullet {
	return &Bullet{}
}

func (b *Bullet) Type() body.Shape {
	return body.Bullet
}
