package strategy

import (
	"fmt"

	"simuduck/pkg/model"
)

func NewRectangleStrategy(left, top, width, height float64) model.ShapeStrategy {
	return &rectangleStrategy{left, top, width, height}
}

type rectangleStrategy struct {
	left, top, width, height float64
}

func (r *rectangleStrategy) Draw(canvas model.Canvas, color model.Color) {
	canvas.SetColor(color)

	canvas.MoveTo(r.left, r.top)
	canvas.LineTo(r.left+r.width, r.top)
	canvas.LineTo(r.left+r.width, r.top+r.height)
	canvas.LineTo(r.left, r.top+r.height)
	canvas.LineTo(r.left, r.top)
}

func (r *rectangleStrategy) Move(dx, dy float64) {
	r.left += dx
	r.top += dy
}

func (r *rectangleStrategy) GetType() model.Type {
	return model.TypeRectangle
}

func (r *rectangleStrategy) GetParams() string {
	return fmt.Sprintf("%.1f %.1f %.1f %.1f", r.left, r.top, r.width, r.height)
}

func (r *rectangleStrategy) Clone() model.ShapeStrategy {
	return &rectangleStrategy{
		left:   r.left,
		top:    r.top,
		width:  r.width,
		height: r.height,
	}
}
