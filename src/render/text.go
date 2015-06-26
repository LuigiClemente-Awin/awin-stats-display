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

	output := make([]string, 4)
	// Loop through each object
	for _, object := range t.objects {
		// Type switch
		switch object := object.(type) {
		case *domain.Arena:
			output[0] = "---------------------------------"
			output[1] = "               |                 "
			output[2] = "               |                 "
			output[3] = "---------------------------------"
		case *domain.Paddle:
			yPosition := int(round(object.GetPosition().Y))
			side := object.GetSide()
			row := output[yPosition]
			middle := row[1:len(row)-1]
			if side == 0 {
				output[yPosition] = strings.Join([]string{"|", middle, " "}, "")
			} else {
				output[yPosition] = strings.Join([]string{" ", middle, "|"}, "")
			}
		default:
			fmt.Println("Dunno", object)
		}
	}
	
	for _, line := range output {
		fmt.Println(line)
	}

}