package main

import (
	"fmt"
	"domain"
	"render"
)

func main() {
	p1 := domain.MakePlayer(1)
	paddle1 := domain.MakePaddle(0)
	p1.SetPaddle(paddle1)
	p1.MovePaddleRandom()

	p2 := domain.MakePlayer(2)
	paddle2 := domain.MakePaddle(1)
	p2.SetPaddle(paddle2)
	p2.MovePaddleRandom()

	arena := domain.MakeArena(10, 4)

	renderer := render.MakeText()

	renderer.AddObject(arena)
	renderer.AddObject(paddle1)
	renderer.AddObject(paddle2)

	renderer.Render()
	
	fmt.Println("Paddle position:", paddle1.GetPosition())
}