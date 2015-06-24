package domain

type Movable interface {
	MoveTo(Pointy)
	GetPosition() Pointy
	IsStatic() bool
	MakeStatic(bool)
}
