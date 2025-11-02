package commands

import "editor/pkg/model"

type InsertParagraphCommand struct {
	Doc  *model.Document
	Pos  int
	Text string
}

func (c *InsertParagraphCommand) Execute() error {
	item := model.DocumentItem{Paragraph: model.NewParagraph(c.Text)}
	c.Doc.Items = append(c.Doc.Items, item)
	if c.Pos < len(c.Doc.Items)-1 {
		copy(c.Doc.Items[c.Pos+1:], c.Doc.Items[c.Pos:])
		c.Doc.Items[c.Pos] = item
	}
	return nil
}

func (c *InsertParagraphCommand) Undo() error {
	c.Doc.Items = append(c.Doc.Items[:c.Pos], c.Doc.Items[c.Pos+1:]...)
	return nil
}

func (c *InsertParagraphCommand) Redo() error {
	item := model.DocumentItem{Paragraph: model.NewParagraph(c.Text)}
	c.Doc.Items = append(c.Doc.Items, item)
	if c.Pos < len(c.Doc.Items)-1 {
		copy(c.Doc.Items[c.Pos+1:], c.Doc.Items[c.Pos:])
		c.Doc.Items[c.Pos] = item
	}
	return nil
}

func (c *InsertParagraphCommand) Type() model.CommandType {
	return model.CommandTypeInsertParagraph
}

func (c *InsertParagraphCommand) CanCoalesceWith(_ model.Command) bool {
	return false
}

func (c *InsertParagraphCommand) Coalesce(_ model.Command) {}
