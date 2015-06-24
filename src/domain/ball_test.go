package domain

import "testing"

type MockPoint struct {
	
}

func TestMoveTo(t *testing.T) {
	ball := new(Ball)
	mockPoint := new(MockPoint)
	// Arrange
	ball.MoveTo(mockPoint)

	// Assert
	ballPosition := ball.GetPosition()
	if ballPosition != mockPoint {
		t.Error("Foo")
	}
}
