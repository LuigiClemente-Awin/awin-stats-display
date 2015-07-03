package main

import (
	//"fmt"
	"domain"
	"render"
	"time"
	"physics"
)

func main() {
	p1 := domain.MakePlayer(1)
	paddle1 := domain.MakePaddle(0)
	p1.SetPaddle(paddle1)
	
	p2 := domain.MakePlayer(2)
	paddle2 := domain.MakePaddle(1)
	p2.SetPaddle(paddle2)

	arena := domain.MakeArena(21, 4)

	ball := domain.MakeBall(2, 2)

	renderer := render.MakeText()
	renderer.AddObject(arena)
	//renderer.AddObject(paddle1)
	//renderer.AddObject(paddle2)
	renderer.AddObject(ball)

	engine := physics.MakeEngine()
	engine.AddObject(ball)

	for {
		print("\033[H\033[2J")
		p1.MovePaddleRandom()
		p2.MovePaddleRandom()
		//ball.MoveTo(domain.Point{8, 2})
		engine.Tick()
		renderer.Render()
		time.Sleep(2 * time.Millisecond)
	}
	
	//fmt.Println("Paddle position:", paddle1.GetPosition())
}