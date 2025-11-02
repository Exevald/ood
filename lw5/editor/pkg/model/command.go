package model

type CommandType int

const (
	CommandTypeInsertParagraph CommandType = iota
	CommandTypeInsertImage
	CommandTypeDeleteItem
	CommandTypeSetTitle
	CommandTypeReplaceText
	CommandTypeResizeImage
)

type Command interface {
	Execute() error
	Undo() error
	Redo() error
	Type() CommandType
	CanCoalesceWith(next Command) bool
	Coalesce(next Command)
	// добавить метод удаления последствий
}
