package shapes

import "pictures/pkg/model"

type Ellipse interface {
	model.Shape

	GetCenter() model.Vertex
	GetHorizontalRadius() float64
	GetVerticalRadius() float64
}

func NewEllipse(center model.Vertex, horizontalRadius, verticalRadius float64, color model.Color) Ellipse {
	return &ellipse{
		center:           center,
		horizontalRadius: horizontalRadius,
		verticalRadius:   verticalRadius,
		color:            color,
	}
}

type ellipse struct {
	center                           model.Vertex
	horizontalRadius, verticalRadius float64
	color                            model.Color
}

func (e *ellipse) Draw(canvas model.Canvas) {
	canvas.SetColor(e.color)
	canvas.DrawEllipse(e.center.X, e.center.Y, e.horizontalRadius, e.verticalRadius)
}

func (e *ellipse) GetColor() model.Color {
	return e.color
}

func (e *ellipse) GetCenter() model.Vertex {
	return e.center
}

func (e *ellipse) GetHorizontalRadius() float64 {
	return e.horizontalRadius
}

func (e *ellipse) GetVerticalRadius() float64 {
	return e.verticalRadius
}
