package tdewolff

import (
	canvaslib "github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"

	"shapes/pkg/model"
)

func NewCanvas(width, height float64, filename string) model.Canvas {
	canvas := canvaslib.New(width, height)
	canvasContext := canvaslib.NewContext(canvas)

	return &canvasLib{
		canvas:        canvas,
		canvasContext: canvasContext,
		currentColor:  model.NewColor(255, 255, 255, 255),
		filename:      filename,
	}
}

type canvasLib struct {
	canvas        *canvaslib.Canvas
	canvasContext *canvaslib.Context
	currentX      float64
	currentY      float64
	currentColor  model.Color
	filename      string
}

func (c *canvasLib) SetColor(color model.Color) {
	c.currentColor = color
}

func (c *canvasLib) MoveTo(x, y float64) {
	c.currentX = x
	c.currentY = y
}

func (c *canvasLib) LineTo(x, y float64) {
	path := &canvaslib.Path{}
	path.MoveTo(c.currentX, c.currentY)
	path.LineTo(x, y)

	c.canvasContext.SetStrokeColor(c.currentColor.ToRGBA())
	c.canvasContext.SetStrokeWidth(1.0)
	c.canvasContext.DrawPath(0, 0, path)

	c.currentX = x
	c.currentY = y
}

func (c *canvasLib) DrawEllipse(cx, cy, rx, ry float64) {
	path := &canvaslib.Path{}
	path.MoveTo(cx+rx, cy)
	path.Arc(rx, ry, 0, 0, 360)

	c.canvasContext.SetFillColor(c.currentColor.ToRGBA())
	c.canvasContext.DrawPath(0, 0, path)
}

func (c *canvasLib) DrawText(left, top, fontSize float64, text string) {
	face := canvaslib.NewFontFamily("arial")
	if err := face.LoadSystemFont("Arial", canvaslib.FontRegular); err != nil {
		face = canvaslib.NewFontFamily("default")
		face.LoadLocalFont("arial", canvaslib.FontRegular)
	}

	font := face.Face(fontSize, c.currentColor.ToRGBA(), canvaslib.FontRegular, canvaslib.FontNormal)

	c.canvasContext.DrawText(left, top, canvaslib.NewTextLine(font, text, canvaslib.Left))
}

func (c *canvasLib) SaveToFile(filename string) error {
	if filename == "" {
		filename = c.filename
	}

	return renderers.Write(filename, c.canvas, canvaslib.DPMM(5.0))
}
