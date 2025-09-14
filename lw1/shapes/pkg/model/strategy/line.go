package strategy

import (
	"fmt"

	"shapes/pkg/model"
)

func NewLineStrategy(x1, y1, x2, y2 float64) model.ShapeStrategy {
	return &lineStrategy{x1, y1, x2, y2}
}

type lineStrategy struct {
	x1, y1, x2, y2 float64
}

func (l *lineStrategy) Draw(canvas model.Canvas, color model.Color) {
	canvas.SetColor(color)

	canvas.MoveTo(l.x1, l.y1)
	canvas.LineTo(l.x2, l.y2)
}

func (l *lineStrategy) Move(dx, dy float64) {
	l.x1 += dx
	l.y1 += dy

	l.x2 += dx
	l.y2 += dy
}

func (l *lineStrategy) GetType() model.Type {
	return model.TypeLine
}

func (l *lineStrategy) GetParams() string {
	return fmt.Sprintf("%.1f %.1f %.1f %.1f", l.x1, l.y1, l.x2, l.y2)
}

func (l *lineStrategy) Clone() model.ShapeStrategy {
	return &lineStrategy{
		x1: l.x1, y1: l.y1,
		x2: l.x2, y2: l.y2,
	}
}
