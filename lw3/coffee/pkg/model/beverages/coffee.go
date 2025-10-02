package beverages

import "coffee/pkg/model"

func NewCoffee() model.Beverage {
	return &coffee{
		beverage: model.NewBeverage("Coffee"),
	}
}

type coffee struct {
	beverage model.Beverage
}

func (c *coffee) GetDescription() string {
	return c.beverage.GetDescription()
}

func (c *coffee) GetCost() float64 {
	return 60
}
