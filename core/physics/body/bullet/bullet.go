package bullet

import (
	"github.com/hueypark/marsettler/core/physics/body"
)

type Bullet struct {
}

func New() *Bullet {
	return &Bullet{}
}

func (b *Bullet) Type() int64 {
	return body.Bullet
}
