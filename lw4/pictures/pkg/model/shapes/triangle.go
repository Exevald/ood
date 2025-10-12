package shapes

import "pictures/pkg/model"

type Triangle interface {
	model.Shape

	GetVertex1() model.Vertex
	GetVertex2() model.Vertex
	GetVertex3() model.Vertex
}

func NewTriangle(v1, v2, v3 model.Vertex, color model.Color) Triangle {
	return &triangle{
		v1: v1, v2: v2, v3: v3,
		color: color,
	}
}

type triangle struct {
	v1, v2, v3 model.Vertex
	color      model.Color
}

func (t *triangle) Draw(canvas model.Canvas) {
	canvas.SetColor(t.color)

	canvas.MoveTo(t.v1.X, t.v1.Y)
	canvas.LineTo(t.v2.X, t.v2.Y)
	canvas.LineTo(t.v3.X, t.v3.Y)
	canvas.LineTo(t.v1.X, t.v1.Y)
}

func (t *triangle) GetColor() model.Color {
	return t.color
}

func (t *triangle) GetVertex1() model.Vertex {
	return t.v1
}

func (t *triangle) GetVertex2() model.Vertex {
	return t.v2
}

func (t *triangle) GetVertex3() model.Vertex {
	return t.v3
}
