package model

type Point struct {
	X, Y float64
}

type Canvas interface {
	SetFillColor(color Color)
	SetLineColor(color Color)
	SetLineWidth(width float64)
	DrawLine(x1, y1, x2, y2 float64)
	DrawEllipse(frame Frame)
	FillEllipse(frame Frame)
	FillPolygon(points []Point)
	SaveToFile(filename string) error
}
