package domain

import "testing"

type PointStub struct {}
func (p *PointStub) GetX() float32 { return 0 }
func (p *PointStub) GetY() float32 { return 0 }

func TestMoveTo(t *testing.T) {
	ball := new(Ball)
	mockPoint := new(PointStub)
	// Arrange
	ball.MoveTo(mockPoint)

	// Assert
	ballPosition := ball.GetPosition()
	if ballPosition != mockPoint {
		t.Error("Positiong not the same as passed in")
	}
}
