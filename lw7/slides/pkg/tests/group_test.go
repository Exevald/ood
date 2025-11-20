package tests

import (
	"testing"

	"slides/pkg/model"
	"slides/pkg/model/shapes"
)

func TestGroupFrame(t *testing.T) {
	g := model.NewGroup()
	rect := shapes.NewRectangle(
		model.FillStyle{Enabled: true, Color: model.NewColor(200, 150, 100, 255)},
		model.LineStyle{Enabled: true, Color: model.Black, Width: 2},
		model.Frame{X: 10, Y: 10, Width: 20, Height: 30},
	)
	ellipse := shapes.NewEllipse(
		model.Frame{X: 50, Y: 60, Width: 40, Height: 20},
		model.FillStyle{Enabled: true, Color: model.NewColor(100, 150, 200, 255)},
		model.LineStyle{Enabled: true, Color: model.Black, Width: 1},
	)
	g.Add(rect)
	g.Add(ellipse)

	frame := g.GetFrame()
	if frame.X != 10 || frame.Y != 10 || frame.Width != 80 || frame.Height != 70 {
		t.Errorf("Unexpected frame: %+v", frame)
	}
}

func TestGroupStyles(t *testing.T) {
	g := model.NewGroup()
	style := model.FillStyle{Enabled: true, Color: model.Red}
	r1 := shapes.NewRectangle(
		style,
		model.LineStyle{Enabled: false},
		model.Frame{X: 0, Y: 0, Width: 10, Height: 10},
	)
	r2 := shapes.NewRectangle(
		style,
		model.LineStyle{Enabled: false},
		model.Frame{X: 0, Y: 0, Width: 10, Height: 10},
	)
	g.Add(r1)
	g.Add(r2)

	if g.GetFillStyle() != style {
		t.Error("Expected common fill style")
	}

	r2.SetFillStyle(model.FillStyle{Enabled: true, Color: model.Green})
	if g.GetFillStyle().Enabled {
		t.Error("Expected undefined fill style after divergence")
	}
}

func TestClone(t *testing.T) {
	g1 := model.NewGroup()
	r := shapes.NewRectangle(
		model.FillStyle{Enabled: true, Color: model.Red},
		model.LineStyle{Enabled: false},
		model.Frame{X: 10, Y: 10, Width: 20, Height: 20},
	)
	g1.Add(r)
	g2 := g1.Clone()

	g2.SetFrame(model.Frame{X: 100, Y: 100, Width: 30, Height: 30})
	originalFrame := r.GetFrame()
	if originalFrame.X == 100 {
		t.Error("Clone affected original")
	}
}
