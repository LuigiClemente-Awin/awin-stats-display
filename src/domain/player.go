package domain

type Player struct {
	number uint
	paddle *Paddle
}

func(p *Player) GetPosition() uint {
	return p.number
}

func (p *Player) SetPaddle(paddle *Paddle) {
	p.paddle = paddle
}

func MakePlayer(number uint) *Player {
	return &Player{number: number}
}