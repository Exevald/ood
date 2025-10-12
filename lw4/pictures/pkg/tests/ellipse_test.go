package tests

import (
	"testing"

	"pictures/pkg/model"
	"pictures/pkg/model/shapes"
)

func TestEllipse(t *testing.T) {
	center := model.Vertex{X: 100, Y: 200}
	rx, ry := 30.0, 20.0
	color := model.Blue

	ellipse := shapes.NewEllipse(center, rx, ry, color)

	if ellipse.GetColor() != color {
		t.Errorf("Color mismatch")
	}
	if ellipse.GetCenter() != center {
		t.Errorf("Center mismatch")
	}
	if ellipse.GetHorizontalRadius() != rx {
		t.Errorf("Horizontal radius mismatch")
	}
	if ellipse.GetVerticalRadius() != ry {
		t.Errorf("Vertical radius mismatch")
	}
}
