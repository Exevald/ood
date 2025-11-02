package commands

import "editor/pkg/model"

func NewDeleteItemCommand(
	doc *model.Document,
	pos int,
	item model.DocumentItem,
) model.Command {
	return &deleteItemCommand{
		doc:  doc,
		pos:  pos,
		item: item,
	}
}

type deleteItemCommand struct {
	doc  *model.Document
	pos  int
	item model.DocumentItem
}

func (c *deleteItemCommand) Execute() error {
	c.item = c.doc.Items[c.pos]
	c.doc.Items = append(c.doc.Items[:c.pos], c.doc.Items[c.pos+1:]...)

	if img := c.item.GetImage(); img != nil {
		c.doc.ItemsToDelete[img.GetPath()] = ""
	}
	return nil
}

func (c *deleteItemCommand) Undo() error {
	c.doc.Items = append(c.doc.Items, c.item)
	if c.pos < len(c.doc.Items)-1 {
		copy(c.doc.Items[c.pos+1:], c.doc.Items[c.pos:])
		c.doc.Items[c.pos] = c.item
	}
	if img := c.item.GetImage(); img != nil {
		delete(c.doc.ItemsToDelete, img.GetPath())
	}
	return nil
}

func (c *deleteItemCommand) Redo() error {
	c.doc.Items = append(c.doc.Items[:c.pos], c.doc.Items[c.pos+1:]...)
	if img := c.item.GetImage(); img != nil {
		c.doc.ItemsToDelete[img.GetPath()] = ""
	}
	return nil
}

func (c *deleteItemCommand) Type() model.CommandType {
	return model.CommandTypeDeleteItem
}

func (c *deleteItemCommand) CanCoalesceWith(_ model.Command) bool {
	return false
}

func (c *deleteItemCommand) Coalesce(_ model.Command) {}
