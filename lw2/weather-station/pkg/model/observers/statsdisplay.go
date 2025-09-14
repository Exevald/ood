package observers

import (
	"fmt"
	"math"

	"weatherstation/pkg/model"
)

func NewStatsDisplay() model.Observer {
	return &statsDisplay{
		minTemperature: math.Inf(1),
		maxTemperature: math.Inf(-1),
	}
}

type statsDisplay struct {
	minTemperature float64
	maxTemperature float64
	accTemperature float64
	countAcc       int
}

func (d *statsDisplay) Update(subjectID string, data model.WeatherInfo) {
	if data.Temperature < d.minTemperature {
		d.minTemperature = data.Temperature
	}
	if data.Temperature > d.maxTemperature {
		d.maxTemperature = data.Temperature
	}
	d.accTemperature += data.Temperature
	d.countAcc++

	fmt.Printf("Max Temp %.2f\n", d.maxTemperature)
	fmt.Printf("Min Temp %.2f\n", d.minTemperature)
	fmt.Printf("Average Temp %.2f\n", d.accTemperature/float64(d.countAcc))
	fmt.Println("----------------")
}
