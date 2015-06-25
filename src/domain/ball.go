package domain

type Ball struct {
	position Point
	isStatic bool
}

func (b *Ball) MoveTo(newPosition Point) {
	b.position = newPosition
}

func (b *Ball) GetPosition() Point {
	return b.position
}

func (b *Ball) IsStatic() bool {
	return b.isStatic
}

func (b *Ball) MakeStatic(static bool) {
	b.isStatic = static
}