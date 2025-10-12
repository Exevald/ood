package tests

import (
	"testing"

	"pictures/pkg/model/shapes"
)

func TestShapeFactory(t *testing.T) {
	factory := shapes.NewShapeFactory()

	tests := []struct {
		input    string
		wantType string
		wantErr  bool
	}{
		{"red rectangle 0 0 10 10", "rectangle", false},
		{"green triangle 0 0 1 1 2 2", "triangle", false},
		{"blue ellipse 5 5 3 2", "ellipse", false},
		{"yellow regular_polygon 10 10 5 6", "regular_polygon", false},
		{"pink circle 0 0 5", "", true},
		{"red rect 0 0", "", true},
		{"invalidcolor rectangle 0 0 1 1", "", true},
	}

	for _, tt := range tests {
		shape, err := factory.CreateShape(tt.input)
		if (err != nil) != tt.wantErr {
			t.Errorf("CreateShape(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			continue
		}
		if tt.wantErr {
			continue
		}

		switch tt.wantType {
		case "rectangle":
			if _, ok := shape.(shapes.Rectangle); !ok {
				t.Errorf("Expected rectangle, got %T", shape)
			}
		case "triangle":
			if _, ok := shape.(shapes.Triangle); !ok {
				t.Errorf("Expected triangle, got %T", shape)
			}
		case "ellipse":
			if _, ok := shape.(shapes.Ellipse); !ok {
				t.Errorf("Expected ellipse, got %T", shape)
			}
		case "regular_polygon":
			if _, ok := shape.(shapes.RegularPolygon); !ok {
				t.Errorf("Expected regularPolygon, got %T", shape)
			}
		}
	}
}
