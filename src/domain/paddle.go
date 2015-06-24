package domain

type Side struct {

}

type Paddle struct {
	side Side
	position Point
}

func (p *Paddle) MoveTo(point Pointy) {
	p.position.y = point.GetY();
}

func (p *Paddle) GetPosition() Pointy {
	return p.position
}