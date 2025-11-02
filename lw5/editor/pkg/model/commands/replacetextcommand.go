package commands

import (
	"fmt"

	"editor/pkg/model"
)

type ReplaceTextCommand struct {
	Doc     *model.Document
	Pos     int
	NewText string
	OldText string
}

func (c *ReplaceTextCommand) Execute() error {
	item := c.Doc.GetItem(c.Pos)
	if item.IsImage() {
		return fmt.Errorf("position %d is not a paragraph", c.Pos)
	}
	para := item.GetParagraph()
	c.OldText = para.GetText()
	para.SetText(c.NewText)
	return nil
}

func (c *ReplaceTextCommand) Undo() error {
	para := c.Doc.GetItem(c.Pos).Paragraph
	para.SetText(c.OldText)
	return nil
}

func (c *ReplaceTextCommand) Redo() error {
	para := c.Doc.GetItem(c.Pos).Paragraph
	para.SetText(c.NewText)
	return nil
}

func (c *ReplaceTextCommand) Type() model.CommandType {
	return model.CommandTypeReplaceText
}

func (c *ReplaceTextCommand) CanCoalesceWith(next model.Command) bool {
	nc, ok := next.(*ReplaceTextCommand)
	return ok && nc.Pos == c.Pos
}

func (c *ReplaceTextCommand) Coalesce(next model.Command) {
	nc := next.(*ReplaceTextCommand)
	c.NewText = nc.NewText
}
