package condiments

import (
	"coffee/pkg/model"
	"fmt"
)

type SyrupType string

const (
	SyrupTypeChocolate SyrupType = "chocolate"
	SyrupTypeMaple     SyrupType = "maple"
)

func NewSyrup(beverage model.Beverage, syrupType SyrupType) model.Condiment {
	return &syrup{
		condiment: model.NewCondiment(beverage),
		syrupType: syrupType,
	}
}

type syrup struct {
	condiment model.Condiment
	syrupType SyrupType
}

func (s *syrup) GetCondimentDescription() string {
	return fmt.Sprintf("%s syrup", s.syrupType)
}

func (s *syrup) GetCondimentCost() float64 {
	return 20
}

func (s *syrup) GetDescription() string {
	return fmt.Sprintf("%s, with %s", s.condiment.GetDescription(), s.GetCondimentDescription())
}

func (s *syrup) GetCost() float64 {
	return s.condiment.GetCost() + s.GetCondimentCost()
}
