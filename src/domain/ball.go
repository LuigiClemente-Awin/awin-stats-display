package domain

type Movable interface {
	MoveTo(Pointy)
	GetPosition() Pointy
	IsStatic() bool
	MakeStatic(bool)
}

type Pointy interface {
	
}

type Point struct {
	x float32
	y float32
}

type Ball struct {
	position Pointy
}

func (b *Ball) MoveTo(newPosition Pointy) {
	b.position = newPosition
}

func (b *Ball) GetPosition() Pointy {
	return b.position
}