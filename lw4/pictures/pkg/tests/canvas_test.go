package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"pictures/pkg/model"
)

func TestMockCanvas_MoveTo(t *testing.T) {
	c := MockCanvas{}
	c.MoveTo(10.5, 20.0)
	output := c.GetOutput()
	assert.Len(t, output, 1)
	assert.Equal(t, "MoveTo(10.5, 20.0)", output[0])
}

func TestMockCanvas_LineTo(t *testing.T) {
	c := MockCanvas{}
	c.LineTo(5, 8)
	output := c.GetOutput()
	assert.Len(t, output, 1)
	assert.Equal(t, "LineTo(5.0, 8.0)", output[0])
}

func TestMockCanvas_DrawEllipse(t *testing.T) {
	c := MockCanvas{}
	c.DrawEllipse(1, 2, 3, 4)
	output := c.GetOutput()
	assert.Len(t, output, 1)
	assert.Equal(t, "DrawEllipse(1.0, 2.0, 3.0, 4.0)", output[0])
}

func TestMockCanvas_DrawText(t *testing.T) {
	c := MockCanvas{}
	c.DrawText(1, 2, 12, "hello")
	output := c.GetOutput()
	assert.Len(t, output, 1)
	assert.Equal(t, "DrawText(1.0, 2.0, 12.0, hello)", output[0])
}

func TestMockCanvas_SaveToFile(t *testing.T) {
	c := MockCanvas{}
	err := c.SaveToFile("filename.txt")
	assert.NoError(t, err)
	output := c.GetOutput()
	assert.Len(t, output, 1)
	assert.Equal(t, "SaveToFile(filename.txt)", output[0])
}

func TestMockCanvas_Clear(t *testing.T) {
	c := MockCanvas{}
	c.SetColor(model.Black)
	assert.NotEmpty(t, c.GetOutput())
	c.Clear()
	assert.Empty(t, c.GetOutput())
}
