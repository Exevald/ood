package picture

import (
	stderrors "errors"

	"github.com/pkg/errors"

	"shapes/pkg/model"
)

var (
	errShapeAlreadyExists = stderrors.New("shape already exists")
	errShapeNotFound      = stderrors.New("shape not found")
)

type Picture interface {
	AddShape(shape model.Shape) error
	DeleteShape(id string) error
	MoveShape(id string, dx, dy float64) error
	MovePicture(dx, dy float64) error
	ChangeColor(id string, color model.Color) error
	ChangeShape(id string, strategy model.ShapeStrategy) error
	DrawShape(id string, canvas model.Canvas) error
	DrawPicture(canvas model.Canvas)
	GetShape(id string) (model.Shape, error)
	ListShapes() []model.Shape
	CloneShape(id, newID string) error
}

func NewPicture() Picture {
	return &picture{
		shapes: make(map[string]model.Shape),
	}
}

type picture struct {
	shapes map[string]model.Shape
}

func (p *picture) AddShape(shape model.Shape) error {
	id := shape.GetID()
	if _, exists := p.shapes[id]; exists {
		return errors.WithStack(errShapeAlreadyExists)
	}
	p.shapes[id] = shape

	return nil
}

func (p *picture) DeleteShape(id string) error {
	if _, exists := p.shapes[id]; !exists {
		return errors.WithStack(errShapeNotFound)
	}
	delete(p.shapes, id)

	return nil
}

func (p *picture) MoveShape(id string, dx, dy float64) error {
	shape, exists := p.shapes[id]
	if !exists {
		return errors.WithStack(errShapeNotFound)
	}
	shape.Move(dx, dy)

	return nil
}

func (p *picture) MovePicture(dx, dy float64) error {
	for _, shape := range p.shapes {
		shape.Move(dx, dy)
	}
	return nil
}

func (p *picture) ChangeColor(id string, color model.Color) error {
	shape, exists := p.shapes[id]
	if !exists {
		return errors.WithStack(errShapeNotFound)
	}
	shape.SetColor(color)

	return nil
}

func (p *picture) ChangeShape(id string, strategy model.ShapeStrategy) error {
	shape, exists := p.shapes[id]
	if !exists {
		return errors.WithStack(errShapeNotFound)
	}

	oldID := shape.GetID()
	oldColor := shape.GetColor()
	newShape := model.NewShape(oldID, oldColor, strategy)
	p.shapes[oldID] = newShape

	return nil
}

func (p *picture) DrawShape(id string, canvas model.Canvas) error {
	shape, exists := p.shapes[id]
	if !exists {
		return errors.WithStack(errShapeNotFound)
	}
	shape.Draw(canvas, shape.GetColor())

	return nil
}

func (p *picture) DrawPicture(canvas model.Canvas) {
	for _, shape := range p.shapes {
		shape.Draw(canvas, shape.GetColor())
	}
}

func (p *picture) GetShape(id string) (model.Shape, error) {
	shape, exists := p.shapes[id]
	if !exists {
		return nil, errors.WithStack(errShapeNotFound)
	}

	return shape, nil
}

func (p *picture) ListShapes() []model.Shape {
	shapes := make([]model.Shape, 0, len(p.shapes))
	for _, shape := range p.shapes {
		shapes = append(shapes, shape)
	}

	return shapes
}

func (p *picture) CloneShape(id, newID string) error {
	if _, exists := p.shapes[newID]; exists {
		return errors.WithStack(errShapeAlreadyExists)
	}

	original, exists := p.shapes[id]
	if !exists {
		return errors.WithStack(errShapeNotFound)
	}
	cloned := original.Clone(newID)
	p.shapes[newID] = cloned

	return nil
}
