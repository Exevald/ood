package tests

import (
	"fmt"

	"pictures/pkg/model"
)

type MockCanvas struct {
	output []string
}

func (m *MockCanvas) SaveToFile(filename string) error {
	m.output = append(m.output, fmt.Sprintf("SaveToFile(%v)", filename))
	return nil
}

func (m *MockCanvas) SetColor(color model.Color) {
	m.output = append(m.output, fmt.Sprintf("SetColor(%v)", color))
}

func (m *MockCanvas) MoveTo(x, y float64) {
	m.output = append(m.output, fmt.Sprintf("MoveTo(%.1f, %.1f)", x, y))
}

func (m *MockCanvas) LineTo(x, y float64) {
	m.output = append(m.output, fmt.Sprintf("LineTo(%.1f, %.1f)", x, y))
}

func (m *MockCanvas) DrawEllipse(cx, cy, rx, ry float64) {
	m.output = append(m.output, fmt.Sprintf("DrawEllipse(%.1f, %.1f, %.1f, %.1f)", cx, cy, rx, ry))
}

func (m *MockCanvas) DrawText(left, top, fontSize float64, text string) {
	m.output = append(m.output, fmt.Sprintf("DrawText(%.1f, %.1f, %.1f, %s)", left, top, fontSize, text))
}

func (m *MockCanvas) GetOutput() []string {
	return m.output
}

func (m *MockCanvas) Clear() {
	m.output = nil
}
