package condiments

import (
	"fmt"

	"coffee/pkg/model"
)

func NewCoconutFlakes(beverage model.Beverage, mass int) model.Condiment {
	if mass < 0 {
		mass = 0
	}

	return &coconutFlake{
		condiment: model.NewCondiment(beverage),
		mass:      mass,
	}
}

type coconutFlake struct {
	condiment model.Condiment
	mass      int
}

func (c *coconutFlake) GetCondimentDescription() string {
	return fmt.Sprintf("coconut flakes %dg", c.mass)
}

func (c *coconutFlake) GetCondimentCost() float64 {
	return float64(1 * c.mass)
}

func (c *coconutFlake) GetDescription() string {
	return fmt.Sprintf("%s, with %s", c.condiment.GetDescription(), c.GetCondimentDescription())
}

func (c *coconutFlake) GetCost() float64 {
	return c.condiment.GetCost() + c.GetCondimentCost()
}
