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

func MakeBall(x uint, y uint) *Ball {
	fx := float32(x)
	fy := float32(y)
	point := Point{fx, fy}
	return &Ball{ position: point }
}