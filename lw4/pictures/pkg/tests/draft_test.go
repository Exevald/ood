package tests

import (
	"testing"

	"pictures/pkg/model"
)

type mockShape struct {
	color model.Color
}

func (m *mockShape) Draw(canvas model.Canvas) {}
func (m *mockShape) GetColor() model.Color    { return m.color }

func TestPictureDraft(t *testing.T) {
	draft := model.NewPictureDraft()
	shape1 := &mockShape{model.Red}
	shape2 := &mockShape{model.Blue}

	draft.AddShape(shape1)
	draft.AddShape(shape2)

	if draft.GetShapesCount() != 2 {
		t.Errorf("Expected 2 shapes, got %d", draft.GetShapesCount())
	}

	if draft.GetShape(0).GetColor() != model.Red {
		t.Errorf("Expected red shape at index 0")
	}
	if draft.GetShape(1).GetColor() != model.Blue {
		t.Errorf("Expected blue shape at index 1")
	}
}
