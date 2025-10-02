package beverages

import (
	"fmt"

	"coffee/pkg/model"
)

type MilkshakeSize string

const (
	MilkshakeSizeSmall  MilkshakeSize = "Small"
	MilkshakeSizeMedium MilkshakeSize = "Medium"
	MilkshakeSizeLarge  MilkshakeSize = "Large"
)

func NewMilkshake(size MilkshakeSize) model.Beverage {
	return &milkshake{
		beverage: model.NewBeverage("Milkshake"),
		size:     size,
	}
}

type milkshake struct {
	beverage model.Beverage
	size     MilkshakeSize
}

func (m *milkshake) GetDescription() string {
	return fmt.Sprintf("%s %s", m.beverage.GetDescription(), m.size)
}

func (m *milkshake) GetCost() float64 {
	switch m.size {
	case MilkshakeSizeSmall:
		return 50
	case MilkshakeSizeMedium:
		return 60
	case MilkshakeSizeLarge:
		return 80
	default:
		return 0
	}
}
