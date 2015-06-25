package domain

type Movable interface {
	MoveTo(Point)
	GetPosition() Point
	IsStatic() bool
}
