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
	
	fmt.Println("Paddle position:", paddle1.GetPosition())
	//p1.MoveTo(domain.Point{});
	//fmt.Println("Paddle position:", p1.GetPosition())
}