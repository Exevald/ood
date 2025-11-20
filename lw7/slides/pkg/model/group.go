package model

import (
	"math"
)

type Group interface {
	Add(child Shape)
	Draw(canvas Canvas)
	SetFrame(frame Frame)
	GetFillStyle() FillStyle
	GetLineStyle() LineStyle
	SetFillStyle(style FillStyle)
	SetLineStyle(style LineStyle)
	Clone() Shape
	GetFrame() Frame
}

func NewGroup() Group {
	return &group{children: make([]Shape, 0)}
}

type group struct {
	children []Shape
}

func (g *group) Add(child Shape) {
	g.children = append(g.children, child)
}

func (g *group) Draw(canvas Canvas) {
	for _, child := range g.children {
		child.Draw(canvas)
	}
}

func (g *group) GetFrame() Frame {
	if len(g.children) == 0 {
		return Frame{}
	}
	minX, minY := math.MaxFloat64, math.MaxFloat64
	maxX, maxY := -math.MaxFloat64, -math.MaxFloat64
	for _, child := range g.children {
		f := child.GetFrame()
		if f.X < minX {
			minX = f.X
		}
		if f.Y < minY {
			minY = f.Y
		}
		if f.X+f.Width > maxX {
			maxX = f.X + f.Width
		}
		if f.Y+f.Height > maxY {
			maxY = f.Y + f.Height
		}
	}
	return Frame{
		X: minX, Y: minY,
		Width: maxX - minX, Height: maxY - minY,
	}
}

func (g *group) SetFrame(frame Frame) {
	if len(g.children) == 0 {
		return
	}
	oldFrame := g.GetFrame()
	if oldFrame.Width == 0 || oldFrame.Height == 0 {
		return
	}
	sx := frame.Width / oldFrame.Width
	sy := frame.Height / oldFrame.Height
	for _, child := range g.children {
		childFrame := child.GetFrame()
		newX := frame.X + (childFrame.X-oldFrame.X)*sx
		newY := frame.Y + (childFrame.Y-oldFrame.Y)*sy
		newW := childFrame.Width * sx
		newH := childFrame.Height * sy
		child.SetFrame(Frame{X: newX, Y: newY, Width: newW, Height: newH})
	}
}

func (g *group) getCommonFillStyle() *FillStyle {
	if len(g.children) == 0 {
		return nil
	}
	ref := g.children[0].GetFillStyle()
	for _, child := range g.children[1:] {
		s := child.GetFillStyle()
		if s.Enabled != ref.Enabled || (s.Enabled && s.Color != ref.Color) {
			return nil
		}
	}
	result := ref
	return &result
}

func (g *group) getCommonLineStyle() *LineStyle {
	if len(g.children) == 0 {
		return nil
	}
	ref := g.children[0].GetLineStyle()
	for _, child := range g.children[1:] {
		s := child.GetLineStyle()
		if s.Enabled != ref.Enabled ||
			(s.Enabled && (s.Color != ref.Color || s.Width != ref.Width)) {
			return nil
		}
	}
	result := ref
	return &result
}

func (g *group) GetFillStyle() FillStyle {
	if common := g.getCommonFillStyle(); common != nil {
		return *common
	}
	return FillStyle{}
}

func (g *group) GetLineStyle() LineStyle {
	if common := g.getCommonLineStyle(); common != nil {
		return *common
	}
	return LineStyle{}
}

func (g *group) SetFillStyle(style FillStyle) {
	for _, child := range g.children {
		child.SetFillStyle(style)
	}
}

func (g *group) SetLineStyle(style LineStyle) {
	for _, child := range g.children {
		child.SetLineStyle(style)
	}
}

func (g *group) Clone() Shape {
	newGroup := NewGroup()
	for _, child := range g.children {
		newGroup.Add(child.Clone())
	}
	return newGroup
}
