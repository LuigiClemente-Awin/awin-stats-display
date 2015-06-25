package domain

type paddle interface {

}

type Player struct {
	number uint
	paddle paddle
}

func(p *Player) GetPosition() uint {
	return p.number
}

func (p *Player) SetPaddle(paddle paddle) {
	p.paddle = paddle
}

func (p *Player) MovePaddleRandom() {

}

func MakePlayer(number uint) *Player {
	return &Player{number: number}
}
