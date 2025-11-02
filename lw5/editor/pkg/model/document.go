package model

import (
	"fmt"
	"html"
	"os"
	"path/filepath"
	"strings"
)

var ImageCounter int64

type DocumentItem struct {
	Paragraph Paragraph
	Image     Image
}

func (di *DocumentItem) GetParagraph() Paragraph {
	return di.Paragraph
}

func (di *DocumentItem) GetImage() Image {
	return di.Image
}

func (di *DocumentItem) IsImage() bool {
	return di.Image != nil
}

func NewDocument(workDir string) *Document {
	imagesDir := filepath.Join(workDir, "images")
	os.MkdirAll(imagesDir, 0755)

	return &Document{
		title:           "Untitled",
		workDir:         workDir,
		commandsHistory: *NewHistory(),
		ItemsToDelete:   make(map[string]string),
	}
}

type Document struct {
	Items           []DocumentItem
	commandsHistory History
	title           string
	workDir         string
	ItemsToDelete   map[string]string
}

func (d *Document) GetTitle() string {
	return d.title
}

func (d *Document) GetItemsCount() int {
	return len(d.Items)
}

func (d *Document) GetItem(index int) DocumentItem {
	return d.Items[index]
}

func (d *Document) GetItems() []DocumentItem {
	return d.Items
}

func (d *Document) SetTitle(title string) {
	d.title = title
}

func (d *Document) GetWorkDir() string {
	return d.workDir
}

func (d *Document) CanUndo() bool {
	return d.commandsHistory.CanUndo()
}

func (d *Document) Undo() error {
	return d.commandsHistory.Undo()
}

func (d *Document) CanRedo() bool {
	return d.commandsHistory.CanRedo()
}

func (d *Document) Redo() error {
	return d.commandsHistory.Redo()
}

func (d *Document) GetHistory() *History {
	return &d.commandsHistory
}

func (d *Document) Save(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	title := html.EscapeString(d.title)
	_, err = file.WriteString(fmt.Sprintf("<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n  <title>%s</title>\n</head>\n<body>\n", title))
	if err != nil {
		return err
	}

	for _, item := range d.Items {
		if para := item.GetParagraph(); para != nil {
			text := html.EscapeString(para.GetText())
			_, err = file.WriteString(fmt.Sprintf("  <p>%s</p>\n", text))
			if err != nil {
				return err
			}
		} else if img := item.GetImage(); img != nil {
			relPath := filepath.ToSlash(strings.TrimPrefix(img.GetPath(), filepath.Dir(path)+string(filepath.Separator)))
			_, err = file.WriteString(fmt.Sprintf("  <img src=\"%s\" width=\"%d\" height=\"%d\" />\n",
				html.EscapeString(relPath), img.GetWidth(), img.GetHeight()))
			if err != nil {
				return err
			}
		}
	}

	_, err = file.WriteString("</body>\n</html>\n")
	return err
}
