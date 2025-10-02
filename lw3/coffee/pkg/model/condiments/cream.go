package condiments

import (
	"fmt"

	"coffee/pkg/model"
)

func NewCream(beverage model.Beverage) model.Condiment {
	return &cream{
		condiment: model.NewCondiment(beverage),
	}
}

type cream struct {
	condiment model.Condiment
}

func (c *cream) GetCondimentDescription() string {
	return "cream"
}

func (c *cream) GetCondimentCost() float64 {
	return 25
}

func (c *cream) GetDescription() string {
	return fmt.Sprintf("%s, with %s", c.condiment.GetDescription(), c.GetCondimentDescription())
}

func (c *cream) GetCost() float64 {
	return c.condiment.GetCost() + c.GetCondimentCost()
}
