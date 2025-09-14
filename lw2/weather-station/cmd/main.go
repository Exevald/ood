package main

import (
	"weatherstation/pkg/model"
	"weatherstation/pkg/model/observers"
)

func main() {
	wd := model.NewWeatherData()

	display := observers.NewDisplay()
	wd.RegisterObserver(display, 0)

	statsDisplay := observers.NewStatsDisplay()
	wd.RegisterObserver(statsDisplay, 1)

	wd.SetMeasurements(3, 0.7, 760)
	wd.SetMeasurements(4, 0.8, 761)
	wd.RemoveObserver(statsDisplay)
	wd.SetMeasurements(10, 0.8, 761)
	wd.SetMeasurements(-10, 0.8, 761)
}
