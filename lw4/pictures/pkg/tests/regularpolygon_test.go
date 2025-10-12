package tests

import (
	"testing"

	"pictures/pkg/model"
	"pictures/pkg/model/shapes"
)

func TestRegularPolygon(t *testing.T) {
	center := model.Vertex{X: 50, Y: 50}
	radius := 25.0
	vertexCount := 5
	color := model.Yellow

	poly := shapes.NewRegularPolygon(center, radius, vertexCount, color)

	if poly.GetColor() != color {
		t.Errorf("Color mismatch")
	}
	if poly.GetCenter() != center {
		t.Errorf("Center mismatch")
	}
	if poly.GetRadius() != radius {
		t.Errorf("Radius mismatch")
	}
	if poly.GetVertexCount() != vertexCount {
		t.Errorf("Vertex count mismatch")
	}
}

func TestInvalidRegularPolygon(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for vertexCount < 3")
		}
	}()
	shapes.NewRegularPolygon(model.Vertex{X: 0, Y: 0}, 10, 2, model.Red)
}
