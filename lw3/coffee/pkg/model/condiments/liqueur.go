package condiments

import (
	"fmt"

	"coffee/pkg/model"
)

type LiqueurType string

const (
	LiqueurTypeNut       LiqueurType = "nut"
	LiqueurTypeChocolate LiqueurType = "chocolate"
)

func NewLiqueur(beverage model.Beverage, liqueurType LiqueurType) model.Condiment {
	return &liqueur{
		condiment:   model.NewCondiment(beverage),
		liqueurType: liqueurType,
	}
}

type liqueur struct {
	condiment   model.Condiment
	liqueurType LiqueurType
}

func (l *liqueur) GetCondimentDescription() string {
	return fmt.Sprintf("%s liqueur", l.liqueurType)
}

func (l *liqueur) GetCondimentCost() float64 {
	return 50
}

func (l *liqueur) GetDescription() string {
	return fmt.Sprintf("%s, with %s", l.condiment.GetDescription(), l.GetCondimentDescription())
}

func (l *liqueur) GetCost() float64 {
	return l.condiment.GetCost() + l.GetCondimentCost()
}
