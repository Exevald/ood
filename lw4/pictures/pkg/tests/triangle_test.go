package tests

import (
	"testing"

	"pictures/pkg/model"
	"pictures/pkg/model/shapes"
)

func TestTriangle(t *testing.T) {
	v1 := model.Vertex{X: 1, Y: 2}
	v2 := model.Vertex{X: 3, Y: 4}
	v3 := model.Vertex{X: 5, Y: 6}
	color := model.Red

	tri := shapes.NewTriangle(v1, v2, v3, color)

	if tri.GetColor() != color {
		t.Errorf("Expected color %v, got %v", color, tri.GetColor())
	}
	if tri.GetVertex1() != v1 {
		t.Errorf("Vertex1 mismatch")
	}
	if tri.GetVertex2() != v2 {
		t.Errorf("Vertex2 mismatch")
	}
	if tri.GetVertex3() != v3 {
		t.Errorf("Vertex3 mismatch")
	}
}
