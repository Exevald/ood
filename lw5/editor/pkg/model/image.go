package model

import (
	"os"

	"github.com/pkg/errors"
)

type Image interface {
	GetPath() string
	GetWidth() int
	GetHeight() int
	Resize(width, height int)
	Remove() error
}

func NewImage(path string, width, height int) Image {
	return &image{
		path:   path,
		width:  width,
		height: height,
	}
}

type image struct {
	path          string
	width, height int
}

func (i *image) GetPath() string {
	return i.path
}

func (i *image) GetWidth() int {
	return i.width
}

func (i *image) GetHeight() int {
	return i.height
}

func (i *image) Resize(width, height int) {
	i.width = width
	i.height = height
}

func (i *image) Remove() error {
	err := os.Remove(i.path)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
