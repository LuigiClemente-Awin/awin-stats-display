package render

import "testing"

func TestMakeText(t *testing.T) {
	MakeText()
}

type MockObject struct {}

func TextAddObject(t *testing.T) {
	text := MakeText()
	object := new(MockObject)
	
	// Act
	text.AddObject(object)

	// Assert
	text.Render()
}