package main

import (
	"bufio"
	"editor/pkg/service"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CommandsHandler struct {
	service *service.DocumentService
}

func NewCommandsHandler(service *service.DocumentService) *CommandsHandler {
	return &CommandsHandler{service: service}
}

func (h *CommandsHandler) Run() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Editor started. Type 'Help' for commands.")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if line == "" {
			continue
		}

		err := h.handleCommand(line)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}

func (h *CommandsHandler) handleCommand(line string) error {
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return nil
	}

	cmd := strings.ToLower(parts[0])
	switch cmd {
	case "help":
		h.printHelp()
	case "list":
		h.list()
	case "settitle":
		return h.handleSetTitle(parts[1:])
	case "insertparagraph":
		return h.handleInsertParagraph(parts[1:])
	case "insertimage":
		return h.handleInsertImage(parts[1:])
	case "deletetitem":
		return h.handleDeleteItem(parts[1:])
	case "replacetext":
		return h.handleReplaceText(parts[1:])
	case "resizeimage":
		return h.handleResizeImage(parts[1:])
	case "undo":
		return h.service.Undo()
	case "redo":
		return h.service.Redo()
	case "save":
		return h.handleSave(parts[1:])
	default:
		return fmt.Errorf("unknown command: %s", parts[0])
	}
	return nil
}

func (h *CommandsHandler) handleSetTitle(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: SetTitle <title>")
	}
	title := strings.Join(args, " ")
	return h.service.SetTitle(title)
}

func (h *CommandsHandler) handleInsertParagraph(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: InsertParagraph <pos|end> <text>")
	}
	posStr, text := args[0], strings.Join(args[1:], " ")
	pos, err := h.parsePosition(posStr, true)
	if err != nil {
		return err
	}
	return h.service.InsertParagraph(text, pos)
}

func (h *CommandsHandler) handleInsertImage(args []string) error {
	if len(args) != 4 {
		return fmt.Errorf("usage: InsertImage <pos|end> <width> <height> <path>")
	}
	posStr, wStr, hStr, path := args[0], args[1], args[2], args[3]
	pos, err := h.parsePosition(posStr, true)
	if err != nil {
		return err
	}
	width, err := strconv.Atoi(wStr)
	if err != nil {
		return fmt.Errorf("invalid width")
	}
	height, err := strconv.Atoi(hStr)
	if err != nil {
		return fmt.Errorf("invalid height")
	}
	return h.service.InsertImage(path, width, height, pos)
}

func (h *CommandsHandler) handleDeleteItem(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: DeleteItem <pos>")
	}
	pos, err := h.parsePosition(args[0], false)
	if err != nil {
		return err
	}
	return h.service.DeleteItem(pos)
}

func (h *CommandsHandler) handleReplaceText(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: ReplaceText <pos> <text>")
	}
	pos, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid position")
	}
	text := strings.Join(args[1:], " ")
	return h.service.ReplaceText(pos, text)
}

func (h *CommandsHandler) handleResizeImage(args []string) error {
	if len(args) != 3 {
		return fmt.Errorf("usage: ResizeImage <pos> <width> <height>")
	}
	pos, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid position")
	}
	width, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("invalid width")
	}
	height, err := strconv.Atoi(args[2])
	if err != nil {
		return fmt.Errorf("invalid height")
	}
	return h.service.ResizeImage(pos, width, height)
}

func (h *CommandsHandler) handleSave(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: Save <path>")
	}
	return h.service.Doc.Save(args[0])
}

func (h *CommandsHandler) parsePosition(posStr string, allowEnd bool) (int, error) {
	if allowEnd && posStr == "end" {
		return len(h.service.Doc.Items), nil
	}
	pos, err := strconv.Atoi(posStr)
	if err != nil {
		return 0, fmt.Errorf("invalid position: %s", posStr)
	}
	return pos, nil
}

func (h *CommandsHandler) list() {
	fmt.Printf("Title: %s\n", h.service.Doc.GetTitle())
	for i, item := range h.service.Doc.Items {
		if para := item.GetParagraph(); para != nil {
			fmt.Printf("%d. Paragraph: %s\n", i+1, para.GetText())
		} else if img := item.GetImage(); img != nil {
			fmt.Printf("%d. Image: %d %d %s\n", i+1, img.GetWidth(), img.GetHeight(), img.GetPath())
		}
	}
}

func (h *CommandsHandler) printHelp() {
	fmt.Println(`Available commands:
		Help
		List
		SetTitle <title>
		InsertParagraph <pos|end> <text>
		InsertImage <pos|end> <width> <height> <path>
		DeleteItem <pos>
		ReplaceText <pos> <text>
		ResizeImage <pos> <width> <height>
		Undo
		Redo
		Save <path>`)
}
