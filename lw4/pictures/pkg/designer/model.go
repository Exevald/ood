package designer

import (
	"bufio"
	"pictures/pkg/model"
	"pictures/pkg/model/shapes"
)

type Designer interface {
	CreateDraft(scanner *bufio.Scanner) (model.PictureDraft, error)
}

func NewDesigner(factory shapes.ShapeFactory) Designer {
	return &designer{shapeFactory: factory}
}

type designer struct {
	shapeFactory shapes.ShapeFactory
}

func (d *designer) CreateDraft(scanner *bufio.Scanner) (model.PictureDraft, error) {
	draft := model.NewPictureDraft()
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		shape, err := d.shapeFactory.CreateShape(line)
		if err != nil {
			return nil, err
		}
		draft.AddShape(shape)
	}
	return draft, nil
}
