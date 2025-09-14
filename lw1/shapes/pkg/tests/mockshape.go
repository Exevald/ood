package tests

import (
	"shapes/pkg/model"
)

type MockShape struct {
	id         string
	color      model.Color
	movedX     float64
	movedY     float64
	shapeType  model.Type
	info       string
	drawCalled bool
}

func NewMockShape(id string, color model.Color, shapeType model.Type) *MockShape {
	return &MockShape{
		id:        id,
		color:     color,
		shapeType: shapeType,
		info:      id + "_info",
	}
}

func (m *MockShape) GetID() string {
	return m.id
}

func (m *MockShape) GetColor() model.Color {
	return m.color
}

func (m *MockShape) SetColor(color model.Color) {
	m.color = color
}

func (m *MockShape) Draw(canvas model.Canvas, color model.Color) {
	m.drawCalled = true
}

func (m *MockShape) Move(dx, dy float64) {
	m.movedX += dx
	m.movedY += dy
}

func (m *MockShape) GetType() model.Type {
	return m.shapeType
}

func (m *MockShape) GetInfo() string {
	return m.info
}

func (m *MockShape) Clone(newID string) model.Shape {
	return &MockShape{
		id:        newID,
		color:     m.color,
		shapeType: m.shapeType,
		info:      m.info + "_clone",
	}
}
