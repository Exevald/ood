package commands

import (
	"editor/pkg/model"
)

type SetTitleCommand struct {
	Doc      *model.Document
	NewTitle string
	OldTitle string
}

func (c *SetTitleCommand) Execute() error {
	c.OldTitle = c.Doc.GetTitle()
	c.Doc.SetTitle(c.NewTitle)
	return nil
}

func (c *SetTitleCommand) Undo() error {
	c.Doc.SetTitle(c.OldTitle)
	return nil
}

func (c *SetTitleCommand) Redo() error {
	c.Doc.SetTitle(c.NewTitle)
	return nil
}

func (c *SetTitleCommand) Type() model.CommandType { return model.CommandTypeSetTitle }

func (c *SetTitleCommand) CanCoalesceWith(next model.Command) bool {
	_, ok := next.(*SetTitleCommand)
	return ok
}

func (c *SetTitleCommand) Coalesce(next model.Command) {
	nc := next.(*SetTitleCommand)
	c.NewTitle = nc.NewTitle
}
