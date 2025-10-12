package shapes

import (
	"math"
	"pictures/pkg/model"
)

type RegularPolygon interface {
	model.Shape
	GetVertexCount() int
	GetCenter() model.Vertex
	GetRadius() float64
}

type regularPolygon struct {
	center      model.Vertex
	radius      float64
	vertexCount int
	color       model.Color
}

func NewRegularPolygon(center model.Vertex, radius float64, vertexCount int, color model.Color) RegularPolygon {
	if vertexCount < 3 {
		panic("vertexCount must be at least 3")
	}
	if radius <= 0 {
		panic("radius must be positive")
	}
	return &regularPolygon{
		center:      center,
		radius:      radius,
		vertexCount: vertexCount,
		color:       color,
	}
}

func (r *regularPolygon) Draw(canvas model.Canvas) {
	canvas.SetColor(r.color)

	angleStep := 2 * math.Pi / float64(r.vertexCount)
	firstX := r.center.X + r.radius*math.Cos(0)
	firstY := r.center.Y + r.radius*math.Sin(0)

	canvas.MoveTo(firstX, firstY)

	for i := 1; i < r.vertexCount; i++ {
		angle := float64(i) * angleStep
		x := r.center.X + r.radius*math.Cos(angle)
		y := r.center.Y + r.radius*math.Sin(angle)
		canvas.LineTo(x, y)
	}

	canvas.LineTo(firstX, firstY)
}

func (r *regularPolygon) GetColor() model.Color {
	return r.color
}

func (r *regularPolygon) GetVertexCount() int {
	return r.vertexCount
}

func (r *regularPolygon) GetCenter() model.Vertex {
	return r.center
}

func (r *regularPolygon) GetRadius() float64 {
	return r.radius
}
