package tdewolff

import (
	canvaslib "github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"

	"slides/pkg/model"
)

func NewCanvas(width, height float64, filename string) model.Canvas {
	canvas := canvaslib.New(width, height)
	canvasContext := canvaslib.NewContext(canvas)

	return &CanvasLib{
		canvas:        canvas,
		canvasContext: canvasContext,
		fillColor:     model.NewColor(255, 255, 255, 255),
		lineColor:     model.NewColor(0, 0, 0, 255),
		lineWidth:     1.0,
		filename:      filename,
	}
}

type CanvasLib struct {
	canvas        *canvaslib.Canvas
	canvasContext *canvaslib.Context
	fillColor     model.Color
	lineColor     model.Color
	lineWidth     float64
	filename      string
}

func (c *CanvasLib) SetFillColor(color model.Color) {
	c.fillColor = color
}

func (c *CanvasLib) SetLineColor(color model.Color) {
	c.lineColor = color
}

func (c *CanvasLib) SetLineWidth(width float64) {
	c.lineWidth = width
}

func (c *CanvasLib) DrawLine(x1, y1, x2, y2 float64) {
	path := &canvaslib.Path{}
	path.MoveTo(x1, y1)
	path.LineTo(x2, y2)

	if c.lineWidth > 0 {
		c.canvasContext.SetStrokeColor(c.lineColor.ToRGBA())
		c.canvasContext.SetStrokeWidth(c.lineWidth)
		c.canvasContext.DrawPath(0, 0, path)
	}
}

func (c *CanvasLib) DrawEllipse(frame model.Frame) {
	cx := frame.X + frame.Width/2
	cy := frame.Y + frame.Height/2
	rx := frame.Width / 2
	ry := frame.Height / 2

	path := &canvaslib.Path{}
	path.MoveTo(cx+rx, cy)
	path.Arc(rx, ry, 0, 0, 360)

	c.canvasContext.SetStrokeColor(c.lineColor.ToRGBA())
	c.canvasContext.SetStrokeWidth(c.lineWidth)
	c.canvasContext.DrawPath(0, 0, path)
}

func (c *CanvasLib) FillEllipse(frame model.Frame) {
	cx := frame.X + frame.Width/2
	cy := frame.Y + frame.Height/2
	rx := frame.Width / 2
	ry := frame.Height / 2

	path := &canvaslib.Path{}
	path.MoveTo(cx+rx, cy)
	path.Arc(rx, ry, 0, 0, 360)

	c.canvasContext.SetFillColor(c.fillColor.ToRGBA())
	c.canvasContext.DrawPath(0, 0, path)
}

func (c *CanvasLib) FillPolygon(points []model.Point) {
	if len(points) == 0 {
		return
	}

	path := &canvaslib.Path{}
	path.MoveTo(points[0].X, points[0].Y)
	for _, pt := range points[1:] {
		path.LineTo(pt.X, pt.Y)
	}
	path.Close()

	c.canvasContext.SetFillColor(c.fillColor.ToRGBA())
	c.canvasContext.DrawPath(0, 0, path)
}

func (c *CanvasLib) SaveToFile(filename string) error {
	if filename == "" {
		filename = c.filename
	}
	return renderers.Write(filename, c.canvas, canvaslib.DPMM(5.0))
}
