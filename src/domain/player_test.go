package domain

import "testing"

func TestMakePlayer(t *testing.T) {
	p := MakePlayer(1)
	if (p.GetPosition() != 1) {
		t.Error("Position isn't returned correctly")
	}
}

func TestSetPaddle(t *testing.T) {
	p := MakePlayer(1)
	paddle := MakePaddle(0)
	p.SetPaddle(paddle)
}

type MockPaddle struct {
	MoveToCallback func(Point)
}

func (m * MockPaddle) MoveTo(p Point) {
	m.MoveToCallback(p)
}

func TestMovePaddleRandom(t *testing.T) {
	// Arrange
	p := MakePlayer(1)
	var called bool
	paddle := &MockPaddle{MoveToCallback: func(p Point) {
		called = true
	}}
	p.SetPaddle(paddle)

	// Act
	p.MovePaddleRandom()
	
	// Assert
	if !called {
		t.Error("Move to method wasn't called on paddle")
	}
}