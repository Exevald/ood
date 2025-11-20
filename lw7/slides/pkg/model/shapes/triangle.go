package shapes

import (
	"slides/pkg/model"
)

func NewTriangle(
	fillStyle model.FillStyle,
	lineStyle model.LineStyle,
	points []model.Point,
	frame model.Frame,
) model.Shape {
	return &triangle{
		fillStyle: fillStyle,
		lineStyle: lineStyle,
		points:    points,
		frame:     frame,
	}
}

type triangle struct {
	fillStyle model.FillStyle
	lineStyle model.LineStyle
	points    []model.Point
	frame     model.Frame
}

func (t *triangle) GetFillStyle() model.FillStyle {
	return t.fillStyle
}

func (t *triangle) SetFillStyle(style model.FillStyle) {
	t.fillStyle = style
}

func (t *triangle) GetLineStyle() model.LineStyle {
	return t.lineStyle
}

func (t *triangle) SetLineStyle(style model.LineStyle) {
	t.lineStyle = style
}

func (t *triangle) Draw(canvas model.Canvas) {
	t.fillTriangle(canvas)
	t.drawTriangleBounds(canvas)
}

func (t *triangle) Clone() model.Shape {
	return &triangle{
		fillStyle: t.fillStyle,
		lineStyle: t.lineStyle,
		points:    t.points,
		frame:     t.frame,
	}
}

func (t *triangle) GetFrame() model.Frame {
	return t.frame
}

func (t *triangle) SetFrame(frame model.Frame) {
	t.frame = frame
}

func (t *triangle) fillTriangle(canvas model.Canvas) {
	if t.fillStyle.Enabled {
		canvas.SetFillColor(t.fillStyle.Color)
		canvas.FillPolygon([]model.Point{
			{t.points[0].X, t.points[0].Y},
			{t.points[1].X, t.points[1].Y},
			{t.points[2].X, t.points[2].Y},
		})
	}
}

func (t *triangle) drawTriangleBounds(canvas model.Canvas) {
	if !t.lineStyle.Enabled {
		return
	}
	canvas.SetLineColor(t.lineStyle.Color)
	canvas.SetLineWidth(t.lineStyle.Width)

	p0 := t.points[0]
	p1 := t.points[1]
	p2 := t.points[2]

	canvas.DrawLine(p0.X, p0.Y, p1.X, p1.Y)
	canvas.DrawLine(p1.X, p1.Y, p2.X, p2.Y)
	canvas.DrawLine(p2.X, p2.Y, p0.X, p0.Y)
}
