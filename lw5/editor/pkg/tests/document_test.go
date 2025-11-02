package tests

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"editor/pkg/model"
	"editor/pkg/service"
)

func TestDocumentService_SetTitle(t *testing.T) {
	doc := model.NewDocument("")
	svc := service.NewDocumentService(doc)
	if err := svc.SetTitle("Test"); err != nil {
		t.Fatalf("SetTitle failed: %v", err)
	}
	if doc.GetTitle() != "Test" {
		t.Errorf("Title not set")
	}
}

func TestDocumentService_InsertParagraph(t *testing.T) {
	doc := model.NewDocument("")
	svc := service.NewDocumentService(doc)
	if err := svc.InsertParagraph("Hello", 0); err != nil {
		t.Fatalf("InsertParagraph failed: %v", err)
	}
	if doc.GetItemsCount() != 1 {
		t.Errorf("Paragraph not inserted")
	}
	item := doc.GetItem(0)
	if item.GetParagraph().GetText() != "Hello" {
		t.Errorf("Wrong text")
	}
}

func TestDocumentService_UndoRedo(t *testing.T) {
	doc := model.NewDocument("")
	documentService := service.NewDocumentService(doc)

	err := documentService.InsertParagraph("A", 0)
	if err != nil {
		t.Error(err.Error())
	}
	err = documentService.InsertParagraph("B", 1)
	if err != nil {
		t.Error(err.Error())
	}

	if doc.GetItemsCount() != 2 {
		t.Fatal("Setup failed")
	}

	if err = documentService.Undo(); err != nil {
		t.Fatalf("Undo failed: %v", err)
	}
	if doc.GetItemsCount() != 1 {
		t.Errorf("Undo didn't remove item")
	}

	if err = documentService.Redo(); err != nil {
		t.Fatalf("Redo failed: %v", err)
	}
	if doc.GetItemsCount() != 2 {
		t.Errorf("Redo didn't restore item")
	}
}

func TestDocument_Save(t *testing.T) {
	tmpDir := t.TempDir()
	doc := model.NewDocument(tmpDir)
	svc := service.NewDocumentService(doc)
	err := svc.SetTitle("Test doc")
	if err != nil {
		t.Error(err.Error())
	}
	err = svc.InsertParagraph("Hello<World", 0)
	if err != nil {
		t.Error(err.Error())
	}

	savePath := filepath.Join(tmpDir, "test.html")
	if err = doc.Save(savePath); err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	content, err := os.ReadFile(savePath)
	if err != nil {
		t.Fatal(err)
	}

	s := string(content)
	if !strings.Contains(s, "<p>Hello&lt;World</p>") {
		t.Error("Paragraph not saved or not escaped properly")
	}
	if !strings.Contains(s, "<title>Test doc</title>") {
		t.Error("Title not saved")
	}
}
