package commands

import (
	"fmt"

	"editor/pkg/model"
)

type ResizeImageCommand struct {
	Doc  *model.Document
	Pos  int
	NewW int
	NewH int
	OldW int
	OldH int
}

func (c *ResizeImageCommand) Execute() error {
	item := c.Doc.GetItem(c.Pos)
	if !item.IsImage() {
		return fmt.Errorf("position %d is not an image", c.Pos)
	}
	img := item.GetImage()
	c.OldW, c.OldH = img.GetWidth(), img.GetHeight()
	img.Resize(c.NewW, c.NewH)
	return nil
}

func (c *ResizeImageCommand) Undo() error {
	img := c.Doc.GetItem(c.Pos).Image
	img.Resize(c.OldW, c.OldH)
	return nil
}

func (c *ResizeImageCommand) Redo() error {
	img := c.Doc.GetItem(c.Pos).Image
	img.Resize(c.NewW, c.NewH)
	return nil
}

func (c *ResizeImageCommand) Type() model.CommandType {
	return model.CommandTypeResizeImage
}

func (c *ResizeImageCommand) CanCoalesceWith(next model.Command) bool {
	nc, ok := next.(*ResizeImageCommand)
	return ok && nc.Pos == c.Pos
}

func (c *ResizeImageCommand) Coalesce(next model.Command) {
	nc := next.(*ResizeImageCommand)
	c.NewW, c.NewH = nc.NewW, nc.NewH
}
