package tests

import (
	"fmt"

	"simuduck/pkg/model"
)

func NewMockCanvas() model.Canvas {
	return &mockCanvas{}
}

type mockCanvas struct {
	output []string
}

func (m *mockCanvas) SaveToFile(filename string) error {
	m.output = append(m.output, fmt.Sprintf("SaveToFile(%v)", filename))
	return nil
}

func (m *mockCanvas) SetColor(color model.Color) {
	m.output = append(m.output, fmt.Sprintf("SetColor(%v)", color))
}

func (m *mockCanvas) MoveTo(x, y float64) {
	m.output = append(m.output, fmt.Sprintf("MoveTo(%.1f, %.1f)", x, y))
}

func (m *mockCanvas) LineTo(x, y float64) {
	m.output = append(m.output, fmt.Sprintf("LineTo(%.1f, %.1f)", x, y))
}

func (m *mockCanvas) DrawEllipse(cx, cy, rx, ry float64) {
	m.output = append(m.output, fmt.Sprintf("DrawEllipse(%.1f, %.1f, %.1f, %.1f)", cx, cy, rx, ry))
}

func (m *mockCanvas) DrawText(left, top, fontSize float64, text string) {
	m.output = append(m.output, fmt.Sprintf("DrawText(%.1f, %.1f, %.1f, %s)", left, top, fontSize, text))
}

func (m *mockCanvas) GetOutput() []string {
	return m.output
}

func (m *mockCanvas) Clear() {
	m.output = nil
}
