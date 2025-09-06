package strategy

import (
	"fmt"

	"simuduck/pkg/model"
)

func NewCircleStrategy(x, y, r float64) model.ShapeStrategy {
	return &circleStrategy{x, y, r}
}

type circleStrategy struct {
	x, y, r float64
}

func (c *circleStrategy) Draw(canvas model.Canvas, color model.Color) {
	canvas.SetColor(color)

	canvas.DrawEllipse(c.x, c.y, c.r, c.r)
}

func (c *circleStrategy) Move(dx, dy float64) {
	c.x += dx
	c.y += dy
}

func (c *circleStrategy) GetType() model.Type {
	return model.TypeCircle
}

func (c *circleStrategy) GetParams() string {
	return fmt.Sprintf("%.1f %.1f %.1f", c.x, c.y, c.r)
}

func (c *circleStrategy) Clone() model.ShapeStrategy {
	return &circleStrategy{
		x: c.x,
		y: c.y,
		r: c.r,
	}
}
