package strategy

import (
	"fmt"

	"simuduck/pkg/model"
)

func NewTriangleStrategy(x1, y1, x2, y2, x3, y3 float64) model.ShapeStrategy {
	return &triangleStrategy{x1, y1, x2, y2, x3, y3}
}

type triangleStrategy struct {
	x1, y1, x2, y2, x3, y3 float64
}

func (t *triangleStrategy) Draw(canvas model.Canvas, color model.Color) {
	canvas.SetColor(color)

	canvas.MoveTo(t.x1, t.y1)
	canvas.LineTo(t.x2, t.y2)
	canvas.LineTo(t.x3, t.y3)
	canvas.LineTo(t.x1, t.y1)
}

func (t *triangleStrategy) Move(dx, dy float64) {
	t.x1 += dx
	t.y1 += dy

	t.x2 += dx
	t.y2 += dy

	t.x3 += dx
	t.y3 += dy
}

func (t *triangleStrategy) GetType() model.Type {
	return model.TypeTriangle
}

func (t *triangleStrategy) GetParams() string {
	return fmt.Sprintf("%.1f %.1f %.1f %.1f %.1f %.1f", t.x1, t.y1, t.x2, t.y2, t.x3, t.y3)
}

func (t *triangleStrategy) Clone() model.ShapeStrategy {
	return &triangleStrategy{
		x1: t.x1, y1: t.y1,
		x2: t.x2, y2: t.y2,
		x3: t.x3, y3: t.y3,
	}
}
