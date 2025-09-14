package main

import (
	"weatherstation/pkg/model"
)

func main() {
	wd := model.NewWeatherData()

	display := model.NewDisplay()
	wd.RegisterObserver(display)

	statsDisplay := model.NewStatsDisplay()
	wd.RegisterObserver(statsDisplay)

	wd.SetMeasurements(3, 0.7, 760)
	wd.SetMeasurements(4, 0.8, 761)
	wd.RemoveObserver(statsDisplay)
	wd.SetMeasurements(10, 0.8, 761)
	wd.SetMeasurements(-10, 0.8, 761)
}
