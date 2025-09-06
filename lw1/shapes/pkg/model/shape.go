package model

import "fmt"

type Type string

const (
	TypeCircle    Type = "circle"
	TypeRectangle Type = "rectangle"
	TypeTriangle  Type = "triangle"
	TypeLine      Type = "line"
	TypeText      Type = "text"
)

type Shape interface {
	GetID() string
	GetColor() Color
	SetColor(Color)
	Draw(canvas Canvas, color Color)
	Move(dx, dy float64)
	GetType() Type
	GetInfo() string
	Clone(newID string) Shape
}

type Canvas interface {
	SetColor(color Color)
	MoveTo(x, y float64)
	LineTo(x, y float64)
	DrawEllipse(cx, cy, rx, ry float64)
	DrawText(left, top, fontSize float64, text string)
	SaveToFile(filename string) error
}

type ShapeStrategy interface {
	Draw(canvas Canvas, color Color)
	Move(dx, dy float64)
	GetType() Type
	GetParams() string
	Clone() ShapeStrategy
}

func NewShape(id string, color Color, shapeStrategy ShapeStrategy) Shape {
	return &shape{
		id:            id,
		color:         color,
		shapeStrategy: shapeStrategy,
	}
}

type shape struct {
	id            string
	color         Color
	shapeStrategy ShapeStrategy
}

func (s *shape) GetID() string {
	return s.id
}

func (s *shape) GetColor() Color {
	return s.color
}

func (s *shape) SetColor(color Color) {
	s.color = color
}

func (s *shape) Draw(canvas Canvas, color Color) {
	s.shapeStrategy.Draw(canvas, color)
}

func (s *shape) Move(dx, dy float64) {
	s.shapeStrategy.Move(dx, dy)
}

func (s *shape) GetType() Type {
	return s.shapeStrategy.GetType()
}

func (s *shape) GetInfo() string {
	return fmt.Sprintf("%s %s %s %s", s.GetType(), s.id, s.color, s.shapeStrategy.GetParams())
}

func (s *shape) Clone(newID string) Shape {
	return &shape{
		id:            newID,
		color:         s.color,
		shapeStrategy: s.shapeStrategy.Clone(),
	}
}
