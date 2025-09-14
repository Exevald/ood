package model

import (
	"fmt"
	"math"
)

func NewDisplay() Observer {
	return &display{}
}

type display struct {
}

func (d *display) Update(data WeatherInfo) {
	fmt.Printf("Current Temp %.2f\n", data.Temperature)
	fmt.Printf("Current Hum %.2f\n", data.Humidity)
	fmt.Printf("Current Pressure %.2f\n", data.Pressure)
	fmt.Println("----------------")
}

func NewStatsDisplay() Observer {
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

func (d *statsDisplay) Update(data WeatherInfo) {
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
