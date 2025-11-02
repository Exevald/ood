package tests

import (
	"testing"

	"editor/pkg/model"
)

func TestParagraph_GetText(t *testing.T) {
	p := model.NewParagraph("Hello World")
	if p.GetText() != "Hello World" {
		t.Errorf("Expected 'Hello World', got %q", p.GetText())
	}
}

func TestParagraph_SetText(t *testing.T) {
	p := model.NewParagraph("Old")
	p.SetText("New")
	if p.GetText() != "New" {
		t.Errorf("Expected 'New', got %q", p.GetText())
	}
}
