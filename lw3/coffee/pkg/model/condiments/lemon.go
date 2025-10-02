package condiments

import (
	"fmt"

	"coffee/pkg/model"
)

func NewLemon(beverage model.Beverage, quantity int) model.Condiment {
	if quantity < 0 {
		quantity = 0
	}

	return &lemon{
		beverage: beverage,
		quantity: quantity,
	}
}

type lemon struct {
	beverage model.Beverage
	quantity int
}

func (l *lemon) GetCondimentDescription() string {
	return fmt.Sprintf("lemon x %d", l.quantity)
}

func (l *lemon) GetCondimentCost() float64 {
	return float64(10 * l.quantity)
}

func (l *lemon) GetDescription() string {
	return fmt.Sprintf("%s, with %s", l.beverage.GetDescription(), l.GetCondimentDescription())
}

func (l *lemon) GetCost() float64 {
	return l.beverage.GetCost() + l.GetCondimentCost()
}
