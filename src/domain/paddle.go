package domain

type Paddle struct {
	position Point
}

func (p *Paddle) MoveTo(point Point) {
	p.position.Y = point.Y;
}

func (p *Paddle) GetPosition() Point {
	return p.position
}