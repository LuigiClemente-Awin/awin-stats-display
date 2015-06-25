package domain

import (
	"math/rand"
	"time"
)

type paddle interface {
	MoveTo(Point)
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
	position := rand.Float32() * 4
	p.paddle.MoveTo(Point{Y: position})
}

func MakePlayer(number uint) *Player {
	rand.Seed( time.Now().UTC().UnixNano())

	return &Player{number: number}
}
