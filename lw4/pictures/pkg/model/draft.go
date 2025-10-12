package model

type PictureDraft interface {
	AddShape(Shape)
	GetShapesCount() int
	GetShape(shapeID int) Shape
}

func NewPictureDraft() PictureDraft {
	return &pictureDraft{
		shapes: make([]Shape, 0),
	}
}

type pictureDraft struct {
	shapes []Shape
}

func (p *pictureDraft) AddShape(s Shape) {
	p.shapes = append(p.shapes, s)
}

func (p *pictureDraft) GetShapesCount() int {
	return len(p.shapes)
}

func (p *pictureDraft) GetShape(shapeID int) Shape {
	return p.shapes[shapeID]
}
