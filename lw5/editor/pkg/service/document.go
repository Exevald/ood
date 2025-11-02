package service

import (
	"fmt"
	"os"

	"editor/pkg/model"
	"editor/pkg/model/commands"
)

type DocumentService struct {
	Doc *model.Document
}

func NewDocumentService(doc *model.Document) *DocumentService {
	return &DocumentService{Doc: doc}
}

func (s *DocumentService) SetTitle(title string) error {
	cmd := &commands.SetTitleCommand{Doc: s.Doc, NewTitle: title}
	if err := cmd.Execute(); err != nil {
		return err
	}
	return s.pushCommand(cmd)
}

func (s *DocumentService) InsertParagraph(text string, pos int) error {
	if pos < 0 || pos > len(s.Doc.Items) {
		return fmt.Errorf("invalid position %d", pos)
	}
	cmd := &commands.InsertParagraphCommand{Doc: s.Doc, Pos: pos, Text: text}
	if err := cmd.Execute(); err != nil {
		return err
	}
	return s.pushCommand(cmd)
}

func (s *DocumentService) InsertImage(srcPath string, width, height, pos int) error {
	if pos < 0 || pos > len(s.Doc.Items) {
		return fmt.Errorf("invalid position %d", pos)
	}
	if width < 1 || width > 10000 || height < 1 || height > 10000 {
		return fmt.Errorf("invalid image size: %dx%d", width, height)
	}
	cmd := &commands.InsertImageCommand{
		Doc:     s.Doc,
		Pos:     pos,
		SrcPath: srcPath,
		Width:   width,
		Height:  height,
	}
	if err := cmd.Execute(); err != nil {
		return err
	}
	return s.pushCommand(cmd)
}

func (s *DocumentService) DeleteItem(pos int) error {
	if pos < 0 || pos >= len(s.Doc.Items) {
		return fmt.Errorf("invalid position %d", pos)
	}
	cmd := commands.NewDeleteItemCommand(s.Doc, pos, s.Doc.Items[pos])
	if err := cmd.Execute(); err != nil {
		return err
	}
	return s.pushCommand(cmd)
}

func (s *DocumentService) ReplaceText(pos int, text string) error {
	if pos < 0 || pos >= len(s.Doc.Items) {
		return fmt.Errorf("invalid position %d", pos)
	}
	item := s.Doc.GetItem(pos)
	if item.IsImage() {
		return fmt.Errorf("position %d is not a paragraph", pos)
	}
	cmd := &commands.ReplaceTextCommand{Doc: s.Doc, Pos: pos, NewText: text}
	if err := cmd.Execute(); err != nil {
		return err
	}
	return s.pushCommand(cmd)
}

func (s *DocumentService) ResizeImage(pos, width, height int) error {
	if pos < 0 || pos >= len(s.Doc.Items) {
		return fmt.Errorf("invalid position %d", pos)
	}
	item := s.Doc.GetItem(pos)
	if !item.IsImage() {
		return fmt.Errorf("position %d is not an image", pos)
	}
	if width < 1 || width > 10000 || height < 1 || height > 10000 {
		return fmt.Errorf("invalid image size: %dx%d", width, height)
	}
	cmd := &commands.ResizeImageCommand{Doc: s.Doc, Pos: pos, NewW: width, NewH: height}
	if err := cmd.Execute(); err != nil {
		return err
	}
	return s.pushCommand(cmd)
}

func (s *DocumentService) Undo() error {
	if !s.Doc.CanUndo() {
		return fmt.Errorf("nothing to undo")
	}
	return s.Doc.Undo()
}

func (s *DocumentService) Redo() error {
	if !s.Doc.CanRedo() {
		return fmt.Errorf("nothing to redo")
	}
	return s.Doc.Redo()
}

func (s *DocumentService) pushCommand(cmd model.Command) error {
	removed := s.Doc.GetHistory().Push(cmd)
	for _, c := range removed {
		if toDelete, ok := c.(interface{ GetToDeletePath() string }); ok {
			if path := toDelete.GetToDeletePath(); path != "" {
				err := os.Remove(path)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
