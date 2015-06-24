package domain

import "testing"

func TestMoveTo(t *testing.T) {
	ball := new(Ball)
	mockPoint := new(Pointy)
	// Arrange
	ball.MoveTo(mockPoint)

	// Assert
	ballPosition := ball.GetPosition()
	if ballPosition != mockPoint {
		t.Error("Foo")
	}
}
