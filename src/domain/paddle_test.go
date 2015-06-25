package domain

import "testing"

func TestPaddle(t *testing.T) {
	// Arrange
	paddle := new(Paddle)
	point := Point{0, 2}
	// Act
	paddle.MoveTo(point)

	// Assert
	position := paddle.GetPosition()
	if position.Y != 2.0 {
		t.Error("Position should equal to 2 and is ", position.Y, position);
	}
}

func TestMakePaddle(t *testing.T) {
	MakePaddle(0)
}