package main

import (
	"fmt"
	"domain"
)

func main() {
	p1 := domain.MakePlayer(1)
	paddle1 := domain.MakePaddle(0)
	p1.SetPaddle(paddle1)
	p1.MovePaddleRandom()

	arena := domain.MakeArena(10, 4)
	_ = arena
	
	fmt.Println("Paddle position:", paddle1.GetPosition())
}