package domain

type Pointy interface {
	GetX() float32
	GetY() float32
}

type Point struct {
	x float32
	y float32
}

func (p Point) GetX() float32 {
	return p.x
}

func (p Point) GetY() float32 {
	return p.y
}
