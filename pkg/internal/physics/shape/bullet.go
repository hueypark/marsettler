package shape

// Bullet represents physics shape bullet.
type Bullet struct {
}

// NewBullet creates new bullet.
func NewBullet() *Bullet {
	return &Bullet{}
}

// Type returns type.
func (b *Bullet) Type() Type {
	return BulletType
}
