package tests

import (
	"sync"

	"github.com/pkg/errors"

	"simuduck/pkg/model"
	"simuduck/pkg/picture"
)

type MockPicture struct {
	picture.Picture

	Shapes map[string]model.Shape

	Calls []string

	ErrAddShape    error
	ErrDeleteShape error
	ErrMoveShape   error
	ErrChangeColor error
	ErrChangeShape error
	ErrDrawShape   error
	ErrGetShape    error
	ErrCloneShape  error

	mu sync.Mutex
}

func NewMockPicture() *MockPicture {
	return &MockPicture{
		Shapes: make(map[string]model.Shape),
		Calls:  []string{},
	}
}

func (m *MockPicture) recordCall(name string) {
	m.mu.Lock()
	m.Calls = append(m.Calls, name)
	m.mu.Unlock()
}

func (m *MockPicture) AddShape(shape model.Shape) error {
	m.recordCall("AddShape")
	if m.ErrAddShape != nil {
		return m.ErrAddShape
	}
	id := shape.GetID()
	if _, exists := m.Shapes[id]; exists {
		return errors.New("shape already exists")
	}
	m.Shapes[id] = shape
	return nil
}

func (m *MockPicture) DeleteShape(id string) error {
	m.recordCall("DeleteShape")
	if m.ErrDeleteShape != nil {
		return m.ErrDeleteShape
	}
	if _, exists := m.Shapes[id]; !exists {
		return errors.New("shape not found")
	}
	delete(m.Shapes, id)
	return nil
}

func (m *MockPicture) MoveShape(id string, dx, dy float64) error {
	m.recordCall("MoveShape")
	if m.ErrMoveShape != nil {
		return m.ErrMoveShape
	}
	shape, exists := m.Shapes[id]
	if !exists {
		return errors.New("shape not found")
	}
	shape.Move(dx, dy)
	return nil
}

func (m *MockPicture) MovePicture(dx, dy float64) error {
	m.recordCall("MovePicture")
	for _, shape := range m.Shapes {
		shape.Move(dx, dy)
	}
	return nil
}

func (m *MockPicture) ChangeColor(id string, color model.Color) error {
	m.recordCall("ChangeColor")
	if m.ErrChangeColor != nil {
		return m.ErrChangeColor
	}
	shape, exists := m.Shapes[id]
	if !exists {
		return errors.New("shape not found")
	}
	shape.SetColor(color)
	return nil
}

func (m *MockPicture) ChangeShape(id string, strategy model.ShapeStrategy) error {
	m.recordCall("ChangeShape")
	if m.ErrChangeShape != nil {
		return m.ErrChangeShape
	}
	oldShape, exists := m.Shapes[id]
	if !exists {
		return errors.New("shape not found")
	}
	newShape := model.NewShape(id, oldShape.GetColor(), strategy)
	m.Shapes[id] = newShape
	return nil
}

func (m *MockPicture) DrawShape(id string, canvas model.Canvas) error {
	m.recordCall("DrawShape")
	if m.ErrDrawShape != nil {
		return m.ErrDrawShape
	}
	shape, exists := m.Shapes[id]
	if !exists {
		return errors.New("shape not found")
	}
	shape.Draw(canvas, shape.GetColor())
	return nil
}

func (m *MockPicture) DrawPicture(canvas model.Canvas) {
	m.recordCall("DrawPicture")
	for _, shape := range m.Shapes {
		shape.Draw(canvas, shape.GetColor())
	}
}

func (m *MockPicture) GetShape(id string) (model.Shape, error) {
	m.recordCall("GetShape")
	if m.ErrGetShape != nil {
		return nil, m.ErrGetShape
	}
	shape, exists := m.Shapes[id]
	if !exists {
		return nil, errors.New("shape not found")
	}
	return shape, nil
}

func (m *MockPicture) ListShapes() []model.Shape {
	m.recordCall("ListShapes")
	shapes := make([]model.Shape, 0, len(m.Shapes))
	for _, shape := range m.Shapes {
		shapes = append(shapes, shape)
	}
	return shapes
}

func (m *MockPicture) CloneShape(id, newID string) error {
	m.recordCall("CloneShape")
	if m.ErrCloneShape != nil {
		return m.ErrCloneShape
	}
	if _, exists := m.Shapes[newID]; exists {
		return errors.New("shape already exists")
	}
	shape, exists := m.Shapes[id]
	if !exists {
		return errors.New("shape not found")
	}
	cloned := shape.Clone(newID)
	m.Shapes[newID] = cloned
	return nil
}
