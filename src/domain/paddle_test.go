package domain

import "testing"

func TestPaddle(t *testing.T) {
	// Arrange
	paddle := new(Paddle)

	// Act
	paddle.MoveTo(2)

	// Assert
	position := paddle.GetPosition()
	if position != 2 {
		t.Error("Position should equal to 2");
	}
}