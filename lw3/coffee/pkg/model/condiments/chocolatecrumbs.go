package condiments

import (
	"fmt"

	"coffee/pkg/model"
)

func NewChocolateCrumbs(beverage model.Beverage, mass int) model.Condiment {
	if mass < 0 {
		mass = 0
	}

	return &chocolateCrumbs{
		condiment: model.NewCondiment(beverage),
		mass:      mass,
	}
}

type chocolateCrumbs struct {
	condiment model.Condiment
	mass      int
}

func (c *chocolateCrumbs) GetCondimentDescription() string {
	return fmt.Sprintf("chocolate crumbs %dg", c.mass)
}

func (c *chocolateCrumbs) GetCondimentCost() float64 {
	return float64(2 * c.mass)
}

func (c *chocolateCrumbs) GetDescription() string {
	return fmt.Sprintf("%s, with %s", c.condiment.GetDescription(), c.GetCondimentDescription())
}

func (c *chocolateCrumbs) GetCost() float64 {
	return c.condiment.GetCost() + c.GetCondimentCost()
}
