package tests

import (
	"testing"

	"pictures/pkg/model"
	"pictures/pkg/model/shapes"
)

func TestRectangle(t *testing.T) {
	lt := model.Vertex{X: 0, Y: 0}
	rb := model.Vertex{X: 10, Y: 20}
	color := model.Green

	rect := shapes.NewRectangle(lt, rb, color)

	if rect.GetColor() != color {
		t.Errorf("Color mismatch")
	}
	if rect.GetLeftTop() != lt {
		t.Errorf("LeftTop mismatch")
	}
	if rect.GetRightBottom() != rb {
		t.Errorf("RightBottom mismatch")
	}
}
