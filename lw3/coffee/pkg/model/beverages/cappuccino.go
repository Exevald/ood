package beverages

import "coffee/pkg/model"

type CappuccinoSize int

const (
	CappuccinoSizeDefault CappuccinoSize = iota
	CappuccinoSizeDouble
)

func NewCappuccino(size CappuccinoSize) model.Beverage {
	return &cappuccino{
		beverage: model.NewBeverage("Cappuccino"),
		size:     size,
	}
}

type cappuccino struct {
	beverage model.Beverage
	size     CappuccinoSize
}

func (c *cappuccino) GetDescription() string {
	return c.beverage.GetDescription()
}

func (c *cappuccino) GetCost() float64 {
	if c.size == CappuccinoSizeDouble {
		return 120
	}
	return 80
}
