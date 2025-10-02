package condiments

import (
	"fmt"

	"coffee/pkg/model"
)

type IceCubeType string

const (
	IceCubeTypeDry   IceCubeType = "dry"
	IceCubeTypeWater IceCubeType = "water"
)

func NewIceCube(beverage model.Beverage, quantity int, cubeType IceCubeType) model.Condiment {
	if quantity < 0 {
		quantity = 0
	}

	return &iceCube{
		beverage: beverage,
		quantity: quantity,
		cubeType: cubeType,
	}
}

type iceCube struct {
	beverage model.Beverage
	quantity int
	cubeType IceCubeType
}

func (i *iceCube) GetCondimentDescription() string {
	return fmt.Sprintf("cube %s x %d", i.cubeType, i.quantity)
}

func (i *iceCube) GetCondimentCost() float64 {
	price := 5.0
	if i.cubeType == IceCubeTypeDry {
		price = 10.0
	}
	return price * float64(i.quantity)
}

func (i *iceCube) GetDescription() string {
	return fmt.Sprintf("%s, with %s", i.beverage.GetDescription(), i.GetCondimentDescription())
}

func (i *iceCube) GetCost() float64 {
	return i.beverage.GetCost() + i.GetCondimentCost()
}
