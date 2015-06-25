package domain

type Player struct {
	number uint
}

func(p *Player) GetPosition() uint {
	return p.number
}

func MakePlayer(number uint) *Player {
	return &Player{number: number}
}