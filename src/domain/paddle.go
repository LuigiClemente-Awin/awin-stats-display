package domain

type Side struct {

}

type Position struct {
	y uint
}

type Paddle struct {
	side Side
	position Position
}

func (p *Paddle) MoveTo(y uint) {
	if (y > 3) {
		panic("Paddle can't go futher than 3")
	}

	if (y < 0) {
		panic("Paddle can't go lower than 0")
	}

	p.position.y = y;
}

func (p *Paddle) GetPosition() uint {
	return p.position.y
}