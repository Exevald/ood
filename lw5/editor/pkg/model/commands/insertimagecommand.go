package commands

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync/atomic"

	"editor/pkg/model"
)

type InsertImageCommand struct {
	Doc     *model.Document
	Pos     int
	SrcPath string
	DstPath string
	Width   int
	Height  int
	Removed bool
}

func (c *InsertImageCommand) Execute() error {
	id := atomic.AddInt64(&model.ImageCounter, 1)
	ext := filepath.Ext(c.SrcPath)
	dstName := fmt.Sprintf("img%d%s", id, ext)
	c.DstPath = filepath.Join(c.Doc.GetWorkDir(), "images", dstName)

	src, err := os.Open(c.SrcPath)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(c.DstPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	img := model.NewImage(c.DstPath, c.Width, c.Height)
	c.Doc.Items = append(c.Doc.Items, model.DocumentItem{Image: img})
	if c.Pos < len(c.Doc.Items)-1 {
		copy(c.Doc.Items[c.Pos+1:], c.Doc.Items[c.Pos:])
		c.Doc.Items[c.Pos] = model.DocumentItem{Image: img}
	}

	return nil
}

func (c *InsertImageCommand) Undo() error {
	c.Doc.ItemsToDelete[c.DstPath] = c.SrcPath
	c.Removed = true

	c.Doc.Items = append(c.Doc.Items[:c.Pos], c.Doc.Items[c.Pos+1:]...)
	return nil
}

func (c *InsertImageCommand) Redo() error {
	img := model.NewImage(c.DstPath, c.Width, c.Height)
	c.Doc.Items = append(c.Doc.Items, model.DocumentItem{Image: img})
	if c.Pos < len(c.Doc.Items)-1 {
		copy(c.Doc.Items[c.Pos+1:], c.Doc.Items[c.Pos:])
		c.Doc.Items[c.Pos] = model.DocumentItem{Image: img}
	}
	delete(c.Doc.ItemsToDelete, c.DstPath)
	c.Removed = false
	return nil
}

func (c *InsertImageCommand) Type() model.CommandType {
	return model.CommandTypeInsertImage
}

func (c *InsertImageCommand) CanCoalesceWith(_ model.Command) bool {
	return false
}

func (c *InsertImageCommand) Coalesce(_ model.Command) {}
