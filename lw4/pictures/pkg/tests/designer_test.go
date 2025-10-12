package tests

import (
	"bufio"
	"strings"
	"testing"

	"pictures/pkg/designer"
	"pictures/pkg/model/shapes"
)

func TestDesigner_CreateDraft(t *testing.T) {
	input := "red rectangle 0 0 10 10\nblue triangle 0 0 1 1 2 2"
	scanner := bufio.NewScanner(strings.NewReader(input))

	factory := shapes.NewShapeFactory()
	d := designer.NewDesigner(factory)

	draft, err := d.CreateDraft(scanner)
	if err != nil {
		t.Fatalf("CreateDraft failed: %v", err)
	}

	if draft.GetShapesCount() != 2 {
		t.Errorf("Expected 2 shapes, got %d", draft.GetShapesCount())
	}
}
