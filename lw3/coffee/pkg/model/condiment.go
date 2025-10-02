package model

type Condiment interface {
	Beverage
	GetCondimentCost() float64
	GetCondimentDescription() string
}

func NewCondiment(beverage Beverage) Condiment {
	return &condiment{
		beverage: beverage,
	}
}

type condiment struct {
	beverage Beverage
}

func (c *condiment) GetDescription() string {
	return c.beverage.GetDescription()
}

func (c *condiment) GetCost() float64 {
	return c.beverage.GetCost()
}

func (c *condiment) GetCondimentCost() float64 {
	panic("GetCondimentCost must be implemented by concrete condiment")
}

func (c *condiment) GetCondimentDescription() string {
	panic("GetCondimentDescription must be implemented by concrete condiment")
}
