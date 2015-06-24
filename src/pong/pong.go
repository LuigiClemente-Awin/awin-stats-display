package main

import (
	"fmt"
	"domain"
)

func main() {
	p1 := new(domain.Paddle)
	fmt.Println("Paddle position:", p1.GetPosition())
	p1.MoveTo(domain.Point{});
	fmt.Println("Paddle position:", p1.GetPosition())
}