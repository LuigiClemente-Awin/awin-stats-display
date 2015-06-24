package domain

type Ball struct {
	position Pointy
}

func (b *Ball) MoveTo(newPosition Pointy) {
	b.position = newPosition
}

func (b *Ball) GetPosition() Pointy {
	return b.position
}