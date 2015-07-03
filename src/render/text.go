package render

import (
	"fmt"
	"domain"
	"math"
	"strings"
)

type Renderer interface {
	Render()
}

type Renderable interface {

}

type Text interface {
	AddObject(Renderable)
	Render()
}

func MakeText() Text {
	return &TextRenderer{}
}

type TextRenderer struct {
	objects []Renderable
}

func (t *TextRenderer) AddObject(r Renderable) {
	t.objects = append(t.objects, r)
}

func (t *TextRenderer) Render() {
	round := func(f float32) float32 {
		return float32(math.Floor(float64(f) + .5))
	}

	output := make([]string, 6)
	// Loop through each object
	for _, object := range t.objects {
		// Type switch
		switch object := object.(type) {
		case *domain.Arena:
			output[0] = "---------------------------------"
			output[1] = "               |                 "
			output[2] = "               |                 "
			output[3] = "               |                 "
			output[4] = "               |                 "
			output[5] = "---------------------------------"
		case *domain.Paddle:
			yPosition := int(round(object.GetPosition().Y)) + 1
			side := object.GetSide()
			row := output[yPosition]
			if side == 0 {
				middle := row[1:len(row)]
				output[yPosition] = strings.Join([]string{"|", middle}, "")
			} else {
				middle := row[0:len(row)-1]
				output[yPosition] = strings.Join([]string{middle, "|"}, "")
			}
		case *domain.Ball:
			position := object.GetPosition()
			y := int(round(position.Y))
			x := int(round(position.X))
			row := output[y]
			left := row[0:x]
			right := row[x+1:len(row)]
			output[y] = strings.Join([]string{left, "o", right}, "")
		default:
			fmt.Println("Dunno", object)
		}
	}
	
	// Clear screen
	for _, line := range output {
		fmt.Println(line)
	}

}