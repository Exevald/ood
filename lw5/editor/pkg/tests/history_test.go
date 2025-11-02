package tests

import (
	"errors"
	"testing"

	"editor/pkg/model"
)

func TestHistory_NewHistory(t *testing.T) {
	h := model.NewHistory()
	if h.GetMaxSize() != 10 {
		t.Errorf("Expected maxSize=10, got %d", h.GetMaxSize())
	}
}

func TestHistory_CanUndoRedo(t *testing.T) {
	h := model.NewHistory()
	if h.CanUndo() || h.CanRedo() {
		t.Error("New history should not allow undo/redo")
	}
}

func TestHistory_UndoRedo(t *testing.T) {
	h := model.NewHistory()
	cmd := &mockCommand{cmdType: model.CommandTypeInsertParagraph}
	h.Push(cmd)

	if !h.CanUndo() {
		t.Error("Should be able to undo after push")
	}
	if err := h.Undo(); err != nil {
		t.Errorf("Undo failed: %v", err)
	}
	if !cmd.undone {
		t.Error("Undo not called")
	}

	if !h.CanRedo() {
		t.Error("Should be able to redo after undo")
	}
	if err := h.Redo(); err != nil {
		t.Errorf("Redo failed: %v", err)
	}
	if !cmd.redoed {
		t.Error("Redo not called")
	}
}

//дою

func TestHistory_Coalescing(t *testing.T) {
	h := model.NewHistory()
	cmd1 := &mockCommand{cmdType: model.CommandTypeSetTitle}
	cmd2 := &mockCommand{cmdType: model.CommandTypeSetTitle}
	removed := h.Push(cmd1)
	if len(removed) != 0 {
		t.Error("No commands should be removed on first push")
	}
	removed = h.Push(cmd2)
	if len(removed) != 0 {
		t.Error("No commands should be removed on coalesce")
	}
	if h.GetCommandsCount() != 1 {
		t.Error("Commands should coalesce into one")
	}
	if !cmd1.coalesceCalled {
		t.Error("Coalesce not called")
	}
}

func TestHistory_NoCoalescing_DifferentTypes(t *testing.T) {
	h := model.NewHistory()
	cmd1 := &mockCommand{cmdType: model.CommandTypeSetTitle}
	cmd2 := &mockCommand{cmdType: model.CommandTypeInsertParagraph}
	h.Push(cmd1)
	h.Push(cmd2)
	if h.GetCommandsCount() != 2 {
		t.Error("Different command types should not coalesce")
	}
}

func TestHistory_ErrorHandling(t *testing.T) {
	h := model.NewHistory()
	errCmd := &mockCommand{
		cmdType: model.CommandTypeInsertParagraph,
		undoErr: errors.New("undo failed"),
		redoErr: errors.New("redo failed"),
	}
	h.Push(errCmd)

	if err := h.Undo(); err == nil || err.Error() != "undo failed" {
		t.Errorf("Undo should return error, got %v", err)
	}
	if err := h.Redo(); err == nil || err.Error() != "redo failed" {
		t.Errorf("Redo should return error, got %v", err)
	}
}
