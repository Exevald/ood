package cli

import (
	"fmt"

	"slides/pkg/model"
)

func NewConsoleCanvas() model.Canvas {
	return &consoleCanvas{}
}

type consoleCanvas struct{}

func (c consoleCanvas) SetFillColor(color model.Color) {
	fmt.Printf("SetFillColor: %s\n", color)
}

func (c consoleCanvas) SetLineColor(color model.Color) {
	fmt.Printf("SetLineColor: %s\n", color)
}

func (c consoleCanvas) SetLineWidth(width float64) {
	fmt.Printf("SetLineWidth: %.2f\n", width)
}

func (c consoleCanvas) DrawLine(x1, y1, x2, y2 float64) {
	fmt.Printf("DrawLine: (%.2f, %.2f) -> (%.2f, %.2f)\n", x1, y1, x2, y2)
}

func (c consoleCanvas) DrawEllipse(frame model.Frame) {
	fmt.Printf("DrawEllipse: %+v\n", frame)
}

func (c consoleCanvas) FillEllipse(frame model.Frame) {
	fmt.Printf("FillEllipse: %+v\n", frame)
}

func (c consoleCanvas) FillPolygon(points []model.Point) {
	fmt.Printf("FillPolygon: %v\n", points)
}

func (c consoleCanvas) SaveToFile(filename string) error {
	fmt.Printf("Saved to file %s", filename)
	return nil
}
