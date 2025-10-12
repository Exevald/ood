package model

type Vertex struct {
	X, Y float64
}

type Shape interface {
	Draw(canvas Canvas)
	GetColor() Color
}

type shape struct{}

type Canvas interface {
	SetColor(color Color)
	MoveTo(x, y float64)
	LineTo(x, y float64)
	DrawEllipse(cx, cy, rx, ry float64)
	DrawText(left, top, fontSize float64, text string)
	SaveToFile(filename string) error
}
