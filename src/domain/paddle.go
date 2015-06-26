package domain

type Paddle struct {
	side uint
	position Point
}

func (p *Paddle) MoveTo(point Point) {
	p.position.Y = point.Y;
}

func (p *Paddle) GetPosition() Point {
	return p.position
}

func (p *Paddle) GetSide() uint {
	return p.side
}

func MakePaddle(side uint) *Paddle {
	return &Paddle{side: side, position: Point{}}
}