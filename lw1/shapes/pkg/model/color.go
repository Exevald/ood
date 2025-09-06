package model

import (
	"fmt"
	stdcolor "image/color"
	"strconv"

	"github.com/pkg/errors"
)

type Color interface {
	ToString() string
	ToRGBA() stdcolor.RGBA
}

func NewColor(r, g, b, a uint8) Color {
	return &color{r, g, b, a}
}

type color struct {
	r, g, b, a uint8
}

func ParseColor(hex string) (Color, error) {
	if len(hex) != 7 || hex[0] != '#' {
		return nil, errors.New("invalid color format")
	}

	r, err := strconv.ParseUint(hex[1:3], 16, 8)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	g, err := strconv.ParseUint(hex[3:5], 16, 8)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	b, err := strconv.ParseUint(hex[5:7], 16, 8)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &color{uint8(r), uint8(g), uint8(b), 255}, nil
}

func (c *color) ToString() string {
	return fmt.Sprintf("#%02X%02X%02X", c.r, c.g, c.b)
}

func (c *color) ToRGBA() stdcolor.RGBA {
	return stdcolor.RGBA{
		R: c.r,
		G: c.g,
		B: c.b,
		A: c.a,
	}
}
