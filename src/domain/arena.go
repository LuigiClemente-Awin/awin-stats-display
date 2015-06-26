package domain

type Arena struct {
	width uint
	height uint
}

func MakeArena(width uint, height uint) *Arena {
	return &Arena{width, height}
}