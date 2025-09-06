package main

import (
	"errors"
	"strconv"
	"strings"

	"simuduck/pkg/model"
	"simuduck/pkg/model/strategy"
	"simuduck/pkg/picture"
)

var errEmptyCommand = errors.New("empty command")

type CommandHandler interface {
	HandleCommand(command string) error
}

func NewCommandHandler(picture picture.Picture, canvas model.Canvas) CommandHandler {
	return &commandHandler{
		picture: picture,
		canvas:  canvas,
	}
}

type commandHandler struct {
	picture picture.Picture
	canvas  model.Canvas
}

func (c *commandHandler) HandleCommand(command string) error {
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return errEmptyCommand
	}

	cmd := parts[0]
	args := parts[1:]

	switch cmd {
	case "AddShape":
		return c.handleAddImage(args)
	case "DeleteShape":
		return c.handleDeleteShape(args)
	case "MoveShape":
		return c.handleMoveShape(args)
	case "MovePicture":
		return c.handleMovePicture(args)
	case "ChangeColor":
		return c.handleChangeColor(args)
	case "ChangeShape":
		return c.handleChangeShape(args)
	case "DrawShape":
		return c.handleDrawShape(args)
	case "DrawPicture":
		return c.handleDrawPicture(args)
	case "GetShape":
		return c.handleGetShape(args)
	case "ListShapes":
		return c.handleListShapes(args)
	case "CloneShape":
		return c.handleCloneShape(args)
	}

	return nil
}

func (c *commandHandler) handleAddImage(args []string) error {
	if len(args) < 4 {
		return errors.New("AddShape requires at least 4 arguments")
	}

	id := args[0]
	color, err := model.ParseColor(args[1])
	if err != nil {
		return err
	}

	shapeType := model.Type(args[2])
	params := args[3:]

	var shapeStrategy model.ShapeStrategy
	switch shapeType {
	case model.TypeCircle:
		if len(args) != 3 {
			return errors.New("circle requires 3 parameters: x y r")
		}
		x, err1 := strconv.ParseFloat(params[0], 64)
		y, err2 := strconv.ParseFloat(params[1], 64)
		r, err3 := strconv.ParseFloat(params[2], 64)
		if err1 != nil || err2 != nil || err3 != nil || r < 0 {
			return errors.New("invalid circle parameters")
		}
		shapeStrategy = strategy.NewCircleStrategy(x, y, r)

	case model.TypeRectangle:
		if len(params) != 4 {
			return errors.New("rectangle requires 4 parameters: left top width height")
		}
		left, err1 := strconv.ParseFloat(params[0], 64)
		top, err2 := strconv.ParseFloat(params[1], 64)
		width, err3 := strconv.ParseFloat(params[2], 64)
		height, err4 := strconv.ParseFloat(params[3], 64)
		if err1 != nil || err2 != nil || err3 != nil || err4 != nil || width < 0 || height < 0 {
			return errors.New("invalid rectangle parameters")
		}
		shapeStrategy = strategy.NewRectangleStrategy(left, top, width, height)

	case model.TypeTriangle:
		if len(params) != 6 {
			return errors.New("triangle requires 6 parameters: x1 y1 x2 y2 x3 y3")
		}
		x1, err1 := strconv.ParseFloat(params[0], 64)
		y1, err2 := strconv.ParseFloat(params[1], 64)
		x2, err3 := strconv.ParseFloat(params[2], 64)
		y2, err4 := strconv.ParseFloat(params[3], 64)
		x3, err5 := strconv.ParseFloat(params[4], 64)
		y3, err6 := strconv.ParseFloat(params[5], 64)
		if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil {
			return errors.New("invalid triangle parameters")
		}
		shapeStrategy = strategy.NewTriangleStrategy(x1, y1, x2, y2, x3, y3)

	case model.TypeLine:
		if len(params) != 4 {
			return errors.New("line requires 4 parameters: x1 y1 x2 y2")
		}
		x1, err1 := strconv.ParseFloat(params[0], 64)
		y1, err2 := strconv.ParseFloat(params[1], 64)
		x2, err3 := strconv.ParseFloat(params[2], 64)
		y2, err4 := strconv.ParseFloat(params[3], 64)
		if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
			return errors.New("invalid line parameters")
		}
		shapeStrategy = strategy.NewLineStrategy(x1, y1, x2, y2)

	case model.TypeText:
		if len(params) < 3 {
			return errors.New("text requires at least 3 parameters: left top fontSize text")
		}
		left, err1 := strconv.ParseFloat(params[0], 64)
		top, err2 := strconv.ParseFloat(params[1], 64)
		fontSize, err3 := strconv.ParseFloat(params[2], 64)
		if err1 != nil || err2 != nil || err3 != nil || fontSize < 0 {
			return errors.New("invalid text parameters")
		}
		text := strings.Join(params[3:], " ")
		shapeStrategy = strategy.NewTextStrategy(left, top, fontSize, text)

	default:
		return errors.New("unsuppotred shape type")
	}

	shape := model.NewShape(id, color, shapeStrategy)
	err = c.picture.AddShape(shape)
	if err != nil {
		return err
	}

	return nil
}

func (c *commandHandler) handleDeleteShape(args []string) error {
	if len(args) != 1 {
		return errors.New("DeleteShape requires 1 argument: id")
	}

	err := c.picture.DeleteShape(args[0])
	if err != nil {
		return err
	}

	return nil
}

func (c *commandHandler) handleMoveShape(args []string) error {
	return nil
}

func (c *commandHandler) handleMovePicture(args []string) error {
	return nil
}

func (c *commandHandler) handleChangeColor(args []string) error {
	return nil
}

func (c *commandHandler) handleChangeShape(args []string) error {
	return nil
}

func (c *commandHandler) handleDrawShape(args []string) error {
	return nil
}

func (c *commandHandler) handleDrawPicture(args []string) error {
	return nil
}

func (c *commandHandler) handleGetShape(args []string) error {
	return nil
}

func (c *commandHandler) handleListShapes(args []string) error {
	return nil
}

func (c *commandHandler) handleCloneShape(args []string) error {
	return nil
}
