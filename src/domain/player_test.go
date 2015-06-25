package domain

import "testing"

func TestMakePlayer(t *testing.T) {
	p := MakePlayer(1)
	if (p.GetPosition() != 1) {
		t.Error("Position isn't returned correctly")
	}
}