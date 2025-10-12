package shapes

import (
	"pictures/pkg/model"
)

type Rectangle interface {
	model.Shape

	GetLeftTop() model.Vertex
	GetRightBottom() model.Vertex
}

func NewRectangle(leftTop, rightBottom model.Vertex, color model.Color) Rectangle {
	return &rectangle{
		leftTop:     leftTop,
		rightBottom: rightBottom,
		color:       color,
	}
}

type rectangle struct {
	leftTop, rightBottom model.Vertex
	color                model.Color
}

func (r *rectangle) Draw(canvas model.Canvas) {
	canvas.SetColor(r.color)

	canvas.MoveTo(r.leftTop.X, r.leftTop.X)
	canvas.LineTo(r.rightBottom.X, r.leftTop.X)
	canvas.LineTo(r.rightBottom.X, r.rightBottom.Y)
	canvas.LineTo(r.leftTop.X, r.rightBottom.Y)
	canvas.LineTo(r.leftTop.X, r.leftTop.Y)
}

func (r *rectangle) GetColor() model.Color {
	return r.color
}

func (r *rectangle) GetLeftTop() model.Vertex {
	return r.leftTop
}

func (r *rectangle) GetRightBottom() model.Vertex {
	return r.rightBottom
}
