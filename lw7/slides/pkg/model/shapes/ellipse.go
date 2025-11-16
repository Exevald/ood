package shapes

import (
	"slides/pkg/model"
)

func NewEllipse(
	frame model.Frame,
	fillStyle model.FillStyle,
	lineStyle model.LineStyle,
) model.Shape {
	return &ellipse{
		frame:     frame,
		fillStyle: fillStyle,
		lineStyle: lineStyle,
	}
}

type ellipse struct {
	frame     model.Frame
	radius    int64
	fillStyle model.FillStyle
	lineStyle model.LineStyle
}

func (e *ellipse) GetFillStyle() model.FillStyle {
	return e.fillStyle
}

func (e *ellipse) SetFillStyle(style model.FillStyle) {
	e.fillStyle = style
}

func (e *ellipse) GetLineStyle() model.LineStyle {
	return e.lineStyle
}

func (e *ellipse) SetLineStyle(style model.LineStyle) {
	e.lineStyle = style
}

func (e *ellipse) Draw(canvas model.Canvas) {
	e.fillEllipse(canvas)
	e.drawEllipseLine(canvas)
}

func (e *ellipse) Clone() model.Shape {
	return &ellipse{
		frame:     e.frame,
		radius:    e.radius,
		fillStyle: e.fillStyle,
		lineStyle: e.lineStyle,
	}
}

func (e *ellipse) GetFrame() model.Frame {
	return e.frame
}

func (e *ellipse) SetFrame(frame model.Frame) {
	e.frame = frame
}

func (e *ellipse) fillEllipse(canvas model.Canvas) {
	if e.fillStyle.Enabled {
		canvas.SetFillColor(e.fillStyle.Color)
		canvas.FillEllipse(e.frame)
	}
}

func (e *ellipse) drawEllipseLine(canvas model.Canvas) {
	if e.lineStyle.Enabled {
		canvas.SetLineColor(e.lineStyle.Color)
		canvas.SetLineWidth(e.lineStyle.Width)
		canvas.DrawEllipse(e.frame)
	}
}
