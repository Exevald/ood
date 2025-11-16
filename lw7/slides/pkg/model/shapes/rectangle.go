package shapes

import (
	"slides/pkg/model"
)

func NewRectangle(
	fillStyle model.FillStyle,
	lineStyle model.LineStyle,
	frame model.Frame,
) model.Shape {
	return &rectangle{
		fillStyle: fillStyle,
		lineStyle: lineStyle,
		frame:     frame,
	}
}

type rectangle struct {
	fillStyle model.FillStyle
	lineStyle model.LineStyle
	frame     model.Frame
}

func (r *rectangle) GetFillStyle() model.FillStyle {
	return r.fillStyle
}

func (r *rectangle) SetFillStyle(style model.FillStyle) {
	r.fillStyle = style
}

func (r *rectangle) GetLineStyle() model.LineStyle {
	return r.lineStyle
}

func (r *rectangle) SetLineStyle(style model.LineStyle) {
	r.lineStyle = style
}

func (r *rectangle) Draw(canvas model.Canvas) {
	r.fillRectangle(canvas)
	r.drawRectangleBounds(canvas)
}

func (r *rectangle) Clone() model.Shape {
	return &rectangle{
		fillStyle: r.fillStyle,
		lineStyle: r.lineStyle,
		frame:     r.frame,
	}
}

func (r *rectangle) GetFrame() model.Frame {
	return r.frame
}

func (r *rectangle) SetFrame(frame model.Frame) {
	r.frame = frame
}

func (r *rectangle) fillRectangle(canvas model.Canvas) {
	if r.fillStyle.Enabled {
		canvas.FillPolygon([]model.Point{
			{r.frame.X, r.frame.Y},
			{r.frame.X + r.frame.Width, r.frame.Y},
			{r.frame.X + r.frame.Width, r.frame.Y + r.frame.Height},
			{r.frame.X, r.frame.Y + r.frame.Height},
		})
	}
}

func (r *rectangle) drawRectangleBounds(canvas model.Canvas) {
	if r.lineStyle.Enabled {
		canvas.DrawLine(r.frame.X, r.frame.Y, r.frame.X+r.frame.Width, r.frame.Y)
		canvas.DrawLine(r.frame.X+r.frame.Width, r.frame.Y, r.frame.X+r.frame.Width, r.frame.Y+r.frame.Height)
		canvas.DrawLine(r.frame.X+r.frame.Width, r.frame.Y+r.frame.Height, r.frame.X, r.frame.Y+r.frame.Height)
		canvas.DrawLine(r.frame.X, r.frame.Y+r.frame.Height, r.frame.X, r.frame.Y)
	}
}
