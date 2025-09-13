package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"simuduck/pkg/model"
)

func TestMockShape_BasicMethods(t *testing.T) {
	col := model.NewColor(10, 20, 30, 40)
	s := NewMockShape("id1", col, model.TypeCircle)

	assert.Equal(t, "id1", s.GetID())
	assert.Equal(t, col, s.GetColor())
	assert.Equal(t, model.TypeCircle, s.GetType())
	assert.NotEmpty(t, s.GetInfo())

	newCol := model.NewColor(50, 60, 70, 80)
	s.SetColor(newCol)
	assert.Equal(t, newCol, s.GetColor())

	s.Move(10, 20)
	assert.Equal(t, 10.0, s.movedX)
	assert.Equal(t, 20.0, s.movedY)

	cloned := s.Clone("newID")
	assert.Equal(t, "newID", cloned.GetID())
	assert.Equal(t, s.GetColor(), cloned.GetColor())
	assert.Equal(t, s.GetType(), cloned.GetType())
}

func TestMockShape_DrawCalled(t *testing.T) {
	s := NewMockShape("id2", model.NewColor(1, 2, 3, 4), model.TypeRectangle)
	assert.False(t, s.drawCalled)

	s.Draw(nil, model.NewColor(0, 0, 0, 0))
	assert.True(t, s.drawCalled)
}

func TestMockShape_IntegrationWithMockPicture(t *testing.T) {
	p := NewMockPicture()
	s := NewMockShape("shape1", model.NewColor(1, 2, 3, 4), model.TypeTriangle)

	err := p.AddShape(s)
	assert.NoError(t, err)

	err = p.MoveShape("shape1", 5, 7)
	assert.NoError(t, err)
	assert.Equal(t, 5.0, s.movedX)
	assert.Equal(t, 7.0, s.movedY)

	err = p.ChangeColor("shape1", model.NewColor(100, 110, 120, 130))
	assert.NoError(t, err)

	var canvas model.Canvas = nil
	err = p.DrawShape("shape1", canvas)
	assert.NoError(t, err)
	assert.True(t, s.drawCalled)
}
