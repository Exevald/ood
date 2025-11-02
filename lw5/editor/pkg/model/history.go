package model

type History struct {
	commands []Command
	undoPos  int
	maxSize  int
}

func NewHistory() *History {
	return &History{maxSize: 10}
}

func (h *History) GetMaxSize() int {
	return h.maxSize
}

func (h *History) GetCommandsCount() int {
	return len(h.commands)
}

func (h *History) CanUndo() bool {
	return h.undoPos > 0
}

func (h *History) CanRedo() bool {
	return h.undoPos < len(h.commands)
}

func (h *History) Undo() error {
	if !h.CanUndo() {
		return nil
	}
	h.undoPos--
	return h.commands[h.undoPos].Undo()
}

func (h *History) Redo() error {
	if !h.CanRedo() {
		return nil
	}
	err := h.commands[h.undoPos].Redo()
	h.undoPos++
	return err
}

func (h *History) Push(cmd Command) []Command {
	h.commands = h.commands[:h.undoPos]

	if h.undoPos > 0 {
		prev := h.commands[h.undoPos-1]
		if prev.Type() == cmd.Type() && prev.CanCoalesceWith(cmd) {
			prev.Coalesce(cmd)
			return nil
		}
	}

	h.commands = append(h.commands, cmd)
	h.undoPos++

	var removedCommands []Command
	if len(h.commands) > h.maxSize {
		removedCommands = h.commands[:1]
		h.commands = h.commands[1:]
		h.undoPos--
	}

	return removedCommands
}
