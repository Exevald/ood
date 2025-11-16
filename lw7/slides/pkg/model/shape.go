package model

type FillStyle struct {
	Enabled bool
	Color   Color
}

type LineStyle struct {
	Enabled bool
	Color   Color
	Width   float64
}

type Frame struct {
	X, Y, Width, Height float64
}

type Shape interface {
	GetFillStyle() FillStyle
	SetFillStyle(style FillStyle)
	GetLineStyle() LineStyle
	SetLineStyle(style LineStyle)

	Draw(canvas Canvas)
	Clone() Shape

	GetFrame() Frame
	SetFrame(frame Frame)
}
