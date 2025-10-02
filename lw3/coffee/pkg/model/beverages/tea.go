package beverages

import (
	"fmt"

	"coffee/pkg/model"
)

type TeaType string

const (
	TeaTypeDaHongPao  = "DaHongPao"
	TeaTypeGood2      = "Good 2.0"
	TeaType7542       = "7542"
	TeaTypeTieguanyin = "Tieguanyin"
)

func NewTea(teaType TeaType) model.Beverage {
	return &tea{
		beverage: model.NewBeverage("Tea"),
		teaType:  teaType,
	}
}

type tea struct {
	beverage model.Beverage
	teaType  TeaType
}

func (t *tea) GetDescription() string {
	return fmt.Sprintf("%s %s", t.beverage.GetDescription(), string(t.teaType))
}

func (t *tea) GetCost() float64 {
	return 30
}
