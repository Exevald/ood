package tests

import "editor/pkg/model"

type mockCommand struct {
	executeErr, undoErr, redoErr error
	executed, undone, redoed     bool
	coalesceCalled               bool
	deletePath                   string
	cmdType                      model.CommandType
}

func (c *mockCommand) Execute() error          { c.executed = true; return c.executeErr }
func (c *mockCommand) Undo() error             { c.undone = true; return c.undoErr }
func (c *mockCommand) Redo() error             { c.redoed = true; return c.redoErr }
func (c *mockCommand) Type() model.CommandType { return c.cmdType }
func (c *mockCommand) CanCoalesceWith(other model.Command) bool {
	_, ok := other.(*mockCommand)
	return ok
}
func (c *mockCommand) Coalesce(_ model.Command) { c.coalesceCalled = true }
func (c *mockCommand) GetToDeletePath() string  { return c.deletePath }
