package strategy

import (
	"fmt"

	"shapes/pkg/model"
)

func NewTextStrategy(left, top, fontSize float64, value string) model.ShapeStrategy {
	return &textStrategy{left, top, fontSize, value}
}

type textStrategy struct {
	left, top, fontSize float64
	value               string
}

func (t *textStrategy) Draw(canvas model.Canvas, color model.Color) {
	canvas.SetColor(color)

	canvas.DrawText(t.left, t.top, t.fontSize, t.value)
}

func (t *textStrategy) Move(dx, dy float64) {
	t.left = dx
	t.top = dy
}

func (t *textStrategy) GetType() model.Type {
	return model.TypeText
}

func (t *textStrategy) GetParams() string {
	return fmt.Sprintf("%.1f %.1f %.1f %s", t.left, t.top, t.fontSize, t.value)
}

func (t *textStrategy) Clone() model.ShapeStrategy {
	return &textStrategy{
		left: t.left, top: t.top,
		fontSize: t.fontSize,
		value:    t.value,
	}
}
