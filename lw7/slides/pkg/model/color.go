package model

import (
	stdcolor "image/color"

	"github.com/pkg/errors"
)

type Color interface {
	ToString() string
	ToRGBA() stdcolor.RGBA
}

type color struct {
	name string
	rgba stdcolor.RGBA
}

var (
	Red    Color = &color{name: "red", rgba: stdcolor.RGBA{R: 255, G: 0, B: 0, A: 255}}
	Green  Color = &color{name: "green", rgba: stdcolor.RGBA{R: 0, G: 255, B: 0, A: 255}}
	Blue   Color = &color{name: "blue", rgba: stdcolor.RGBA{R: 0, G: 0, B: 255, A: 255}}
	Yellow Color = &color{name: "yellow", rgba: stdcolor.RGBA{R: 255, G: 255, B: 0, A: 255}}
	Pink   Color = &color{name: "pink", rgba: stdcolor.RGBA{R: 255, G: 192, B: 203, A: 255}}
	Black  Color = &color{name: "black", rgba: stdcolor.RGBA{R: 0, G: 0, B: 0, A: 255}}
)

func NewColor(r, g, b, a uint8) Color {
	return &color{
		rgba: stdcolor.RGBA{R: r, G: g, B: b, A: a},
	}
}

func ParseColor(s string) (Color, error) {
	switch s {
	case "red":
		return Red, nil
	case "green":
		return Green, nil
	case "blue":
		return Blue, nil
	case "yellow":
		return Yellow, nil
	case "pink":
		return Pink, nil
	case "black":
		return Black, nil
	default:
		return nil, errors.Errorf("unknown color: %s (expected one of: red, green, blue, yellow, pink, black)", s)
	}
}

func (c *color) ToString() string {
	return c.name
}

func (c *color) ToRGBA() stdcolor.RGBA {
	return c.rgba
}

func (c *color) String() string {
	return c.ToString()
}
