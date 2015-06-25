package domain

import "testing"

func TestMoveTo(t *testing.T) {
	ball := new(Ball)
	point := Point{1, 2}
	// Arrange
	ball.MoveTo(point)

	// Assert
	ballPosition := ball.GetPosition()
	if ballPosition != point {
		t.Error("Positiong not the same as passed in")
	}
}
