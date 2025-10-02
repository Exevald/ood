package condiments

import (
	"fmt"

	"coffee/pkg/model"
)

func NewCinnamon(beverage model.Beverage) model.Condiment {
	return &cinnamon{
		condiment: model.NewCondiment(beverage),
	}
}

type cinnamon struct {
	condiment model.Condiment
}

func (c *cinnamon) GetCondimentDescription() string {
	return "cinnamon"
}

func (c *cinnamon) GetCondimentCost() float64 {
	return 20
}

func (c *cinnamon) GetDescription() string {
	return fmt.Sprintf("%s, with %s", c.condiment.GetDescription(), c.GetCondimentDescription())
}

func (c *cinnamon) GetCost() float64 {
	return c.condiment.GetCost() + c.GetCondimentCost()
}
