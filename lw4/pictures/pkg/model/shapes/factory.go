package shapes

import (
	"fmt"
	"strconv"
	"strings"

	"pictures/pkg/model"
)

type ShapeFactory interface {
	CreateShape(description string) (model.Shape, error)
}

func NewShapeFactory() ShapeFactory {
	return &shapeFactory{}
}

type shapeFactory struct{}

func (s *shapeFactory) CreateShape(description string) (model.Shape, error) {
	parts := strings.Fields(description)
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid count of description params: %s", description)
	}

	colorStr := parts[0]
	shapeColor, err := model.ParseColor(colorStr)
	if err != nil {
		return nil, err
	}

	shapeType := parts[1]
	switch shapeType {
	case "rectangle":
		if len(parts) != 6 {
			return nil, fmt.Errorf("invalid count of rectangle shape params: %s", description)
		}
		x1, _ := strconv.ParseFloat(parts[2], 64)
		y1, _ := strconv.ParseFloat(parts[3], 64)
		x2, _ := strconv.ParseFloat(parts[4], 64)
		y2, _ := strconv.ParseFloat(parts[5], 64)
		return NewRectangle(model.Vertex{X: x1, Y: y1}, model.Vertex{X: x2, Y: y2}, shapeColor), nil

	case "triangle":
		if len(parts) != 8 {
			return nil, fmt.Errorf("invalid count of triangle shape params: %s", description)
		}
		p1 := model.Vertex{X: parseFloat(parts[2]), Y: parseFloat(parts[3])}
		p2 := model.Vertex{X: parseFloat(parts[4]), Y: parseFloat(parts[5])}
		p3 := model.Vertex{X: parseFloat(parts[6]), Y: parseFloat(parts[7])}
		return NewTriangle(p1, p2, p3, shapeColor), nil

	case "ellipse":
		if len(parts) != 6 {
			return nil, fmt.Errorf("invalid count of ellipse shape params: %s", description)
		}
		cx := parseFloat(parts[2])
		cy := parseFloat(parts[3])
		rx := parseFloat(parts[4])
		ry := parseFloat(parts[5])
		return NewEllipse(model.Vertex{X: cx, Y: cy}, rx, ry, shapeColor), nil

	case "regular_polygon":
		if len(parts) != 6 {
			return nil, fmt.Errorf("invalid count of regular polygon shape params: %s", description)
		}
		cx := parseFloat(parts[2])
		cy := parseFloat(parts[3])
		radius := parseFloat(parts[4])
		vertexCount := int(parseFloat(parts[5]))
		return NewRegularPolygon(model.Vertex{X: cx, Y: cy}, radius, vertexCount, shapeColor), nil
	default:
		return nil, fmt.Errorf("invalid shape type: %s", shapeType)
	}
}

func parseFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}
