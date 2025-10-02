package model

type Beverage interface {
	GetDescription() string
	GetCost() float64
}

func NewBeverage(description string) Beverage {
	return &beverage{
		description: description,
	}
}

type beverage struct {
	description string
}

func (b *beverage) GetDescription() string {
	return b.description
}

func (b *beverage) GetCost() float64 {
	panic("GetCost must be implemented by concrete beverage")
}
