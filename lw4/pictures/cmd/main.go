package main

import (
	"bufio"
	"fmt"
	"os"

	"pictures/pkg/canvas/tdewolff"
	"pictures/pkg/designer"
	"pictures/pkg/model/shapes"
	"pictures/pkg/painter"
)

func main() {
	factory := shapes.NewShapeFactory()
	d := designer.NewDesigner(factory)

	scanner := bufio.NewScanner(os.Stdin)
	draft, err := d.CreateDraft(scanner)
	if err != nil {
		panic(fmt.Errorf("error creating draft: %v", err))
	}

	canvas := tdewolff.NewCanvas(100, 100, "pictures.png")

	p := painter.NewPainter()
	p.DrawPicture(draft, canvas)

	if err := canvas.SaveToFile(""); err != nil {
		panic(fmt.Errorf("error file saving: %v", err))
	}
}
