package condiments

import (
	"fmt"

	"coffee/pkg/model"
)

func NewChocolateSlice(beverage model.Beverage, count int) model.Condiment {
	if count > 5 {
		count = 5
	}
	if count < 0 {
		count = 0
	}

	return &chocolateSlice{
		condiment: model.NewCondiment(beverage),
		count:     count,
	}
}

type chocolateSlice struct {
	condiment model.Condiment
	count     int
}

func (c *chocolateSlice) GetCondimentDescription() string {
	return fmt.Sprintf("cholocate slice x %d", c.count)
}

func (c *chocolateSlice) GetDescription() string {
	return fmt.Sprintf("%s, with %s", c.condiment.GetDescription(), c.GetCondimentDescription())
}

func (c *chocolateSlice) GetCost() float64 {
	return c.condiment.GetCost() + c.GetCondimentCost()
}

func (c *chocolateSlice) GetCondimentCost() float64 {
	return float64(c.count * 10)
}
