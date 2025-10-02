package beverages

import "coffee/pkg/model"

type LatteSize int

const (
	LatteSizeDefault LatteSize = iota
	LatteSizeDouble
)

func NewLatte(size LatteSize) model.Beverage {
	return &latte{
		beverage: model.NewBeverage("Latte"),
		size:     size,
	}
}

type latte struct {
	beverage model.Beverage
	size     LatteSize
}

func (l *latte) GetDescription() string {
	return l.beverage.GetDescription()
}

func (l *latte) GetCost() float64 {
	if l.size == LatteSizeDouble {
		return 130
	}
	return 90
}
