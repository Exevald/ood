package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"simuduck/pkg/model"
	"simuduck/pkg/picture"
)

func TestPicture_AddDelete(t *testing.T) {
	p := picture.NewPicture()

	s := newTestShape("id1")
	err := p.AddShape(s)
	assert.NoError(t, err)

	err = p.AddShape(s)
	assert.Error(t, err)

	err = p.DeleteShape("id1")
	assert.NoError(t, err)

	err = p.DeleteShape("id1")
	assert.Error(t, err)
}

func TestPicture_MoveChangeColor(t *testing.T) {
	p := picture.NewPicture()
	s := newTestShape("id1")
	p.AddShape(s)

	err := p.MoveShape("id1", 10, 20)
	assert.NoError(t, err)

	newColor := model.NewColor(10, 20, 30, 40)
	err = p.ChangeColor("id1", newColor)
	assert.NoError(t, err)
	assert.Equal(t, newColor, s.GetColor())
}

func TestPicture_DrawShapeAndPicture(t *testing.T) {
	p := picture.NewPicture()
	s1 := newTestShape("id1")
	s2 := newTestShape("id2")
	p.AddShape(s1)
	p.AddShape(s2)

	var canvas model.Canvas = nil

	err := p.DrawShape("id1", canvas)
	assert.NoError(t, err)

	p.DrawPicture(canvas)
}

func TestPicture_CloneShape(t *testing.T) {
	p := picture.NewPicture()
	s := newTestShape("id1")
	p.AddShape(s)

	err := p.CloneShape("id1", "id2")
	assert.NoError(t, err)

	cloned, err := p.GetShape("id2")
	assert.NoError(t, err)
	assert.Equal(t, "id2", cloned.GetID())

	err = p.CloneShape("id1", "id2")
	assert.Error(t, err)
}

func newTestShape(id string) model.Shape {
	color := model.NewColor(1, 2, 3, 4)
	strategy := &mockShapeStrategy{}
	return model.NewShape(id, color, strategy)
}

type mockShapeStrategy struct{}

func (m *mockShapeStrategy) Draw(canvas model.Canvas, color model.Color) {}

func (m *mockShapeStrategy) Move(dx, dy float64) {}

func (m *mockShapeStrategy) GetType() model.Type {
	return model.TypeCircle
}

func (m *mockShapeStrategy) GetParams() string {
	return "mockParams"
}

func (m *mockShapeStrategy) Clone() model.ShapeStrategy {
	return &mockShapeStrategy{}
}
