package physics

import (
	"domain"
	"time"
	"math"
	"fmt"
)

type Vector struct {
	X float32
	Y float32
	Velocity float64 // in cells per second?
	Heading float64 // in degrees 0 to 360
	Body Body
}

type Body interface {
	GetPosition() domain.Point
	IsStatic() bool
	MoveTo(domain.Point)
}

type Engine struct {
	bodies []Body
	vectors []Vector
	lastTick time.Time
}

func (e *Engine) AddObject(body Body) {
	e.bodies = append(e.bodies, body)
	if !body.IsStatic() {
		vector := Vector{
			X: body.GetPosition().X,
			Y: body.GetPosition().Y,
			Velocity: 0.1,
			Heading: 19,
			Body: body,
		}
		e.vectors = append(e.vectors, vector)
	}
}

func (e *Engine) Tick() {
	for key, vector := range e.vectors {
		x := float64(vector.Body.GetPosition().X)
		y := float64(vector.Body.GetPosition().Y)

		radians := vector.Heading * (math.Pi / 180)
		newx := math.Cos(radians) * vector.Velocity
		newy := math.Sin(radians) * vector.Velocity
		fmt.Println("X", x, "Y", y, "NewX", newx, "NewY", newy)
		fx := float32(x+newx)
		fy := float32(y+newy)
		point := domain.Point{fx, fy}

		if point.X < 0 {
			// Hit the left side
			q := 90 - e.vectors[key].Heading
			e.vectors[key].Heading += 2 * q
			continue
		}

		if point.X > 31 {
			// Hit the right side
			q := 90 - e.vectors[key].Heading
			e.vectors[key].Heading += 2 * q
			continue
		}
		
		if point.Y < 0 {
			// Hit the roof
			e.vectors[key].Heading = (-1) * e.vectors[key].Heading
			continue
		}

		if point.Y > 5 {
			// Hit the floor
			//a := e.vectors[key].Heading
			e.vectors[key].Heading = (-1) * e.vectors[key].Heading
			continue
		}

		radians = vector.Heading * (math.Pi / 180)
		newx = math.Cos(radians) * vector.Velocity
		newy = math.Sin(radians) * vector.Velocity
		fmt.Println("X", x, "Y", y, "NewX", newx, "NewY", newy)
		fx = float32(x+newx)
		fy = float32(y+newy)
		point = domain.Point{fx, fy}
		vector.Body.MoveTo(point)

	}
}

func MakeEngine() *Engine {
	return &Engine{}
}
