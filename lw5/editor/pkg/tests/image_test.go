package tests

import (
	"os"
	"path/filepath"
	"testing"

	"editor/pkg/model"
)

func TestImage_Getters(t *testing.T) {
	img := model.NewImage("/test/image.png", 800, 600)
	if img.GetPath() != "/test/image.png" {
		t.Errorf("Unexpected path: %s", img.GetPath())
	}
	if img.GetWidth() != 800 {
		t.Errorf("Unexpected width: %d", img.GetWidth())
	}
	if img.GetHeight() != 600 {
		t.Errorf("Unexpected height: %d", img.GetHeight())
	}
}

func TestImage_Resize(t *testing.T) {
	img := model.NewImage("img.jpg", 100, 200)
	img.Resize(300, 400)
	if img.GetWidth() != 300 || img.GetHeight() != 400 {
		t.Errorf("Resize failed: got %dx%d", img.GetWidth(), img.GetHeight())
	}
}

func TestImage_Remove(t *testing.T) {
	tmpFile := filepath.Join(t.TempDir(), "test.png")
	err := os.WriteFile(tmpFile, []byte("fake image content"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	img := model.NewImage(tmpFile, 100, 100)
	if err := img.Remove(); err != nil {
		t.Fatalf("Remove failed: %v", err)
	}

	if _, err := os.Stat(tmpFile); !os.IsNotExist(err) {
		t.Error("File was not deleted")
	}
}
