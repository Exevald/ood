package main

import (
	"log"
	"slides/pkg/canvas/cli"
	"slides/pkg/canvas/tdewolff"
	"slides/pkg/model"
	"slides/pkg/model/shapes"
)

func main() {
	slide := model.NewGroup()

	house := model.NewGroup()
	body := shapes.NewRectangle(
		model.FillStyle{Enabled: true, Color: model.NewColor(200, 150, 100, 255)},
		model.LineStyle{Enabled: true, Color: model.Black, Width: 2},
		model.Frame{X: 100, Y: 200, Width: 200, Height: 150},
	)
	minX, minY := 100.0, 100.0
	maxX, maxY := 300.0, 200.0
	frame := model.Frame{
		X:      minX,
		Y:      minY,
		Width:  maxX - minX,
		Height: maxY - minY,
	}
	roof := shapes.NewTriangle(
		model.FillStyle{Enabled: true, Color: model.NewColor(200, 50, 50, 255)},
		model.LineStyle{Enabled: true, Color: model.Black, Width: 2},
		[]model.Point{
			{X: 100, Y: 200},
			{X: 200, Y: 100},
			{X: 300, Y: 200},
		},
		frame,
	)
	house.Add(body)
	house.Add(roof)

	door := shapes.NewRectangle(
		model.FillStyle{Enabled: true, Color: model.NewColor(100, 80, 50, 255)},
		model.LineStyle{Enabled: true, Color: model.Black, Width: 1.5},
		model.Frame{X: 180, Y: 280, Width: 40, Height: 70},
	)
	house.Add(door)

	window := shapes.NewEllipse(
		model.Frame{X: 240, Y: 240, Width: 30, Height: 30},
		model.FillStyle{Enabled: true, Color: model.NewColor(180, 220, 255, 255)},
		model.LineStyle{Enabled: true, Color: model.Black, Width: 1},
	)
	house.Add(window)

	slide.Add(house)

	sun := shapes.NewEllipse(
		model.Frame{X: 400, Y: 50, Width: 60, Height: 60},
		model.FillStyle{Enabled: true, Color: model.Yellow},
		model.LineStyle{Enabled: false},
	)
	slide.Add(sun)

	slideClone := slide.Clone()
	if clone, ok := slideClone.(*model.Group); ok {
		clone.SetFrame(model.Frame{X: 0, Y: 0, Width: 600, Height: 400}) // move/scale if needed
	}

	cliCanvas := cli.NewConsoleCanvas()
	slide.Draw(cliCanvas)

	canvas := tdewolff.NewCanvas(800, 600, "slides")
	slide.Draw(canvas)
	err := canvas.SaveToFile("slide.png")
	if err != nil {
		log.Fatal(err)
	}
}
