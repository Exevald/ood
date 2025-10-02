package main

import (
	"weatherstation/pkg/model"
	"weatherstation/pkg/model/observers"
)

func main() {
	bus := model.NewEventBus()
	wd := model.NewIndoorWeatherData("station", bus)

	display := observers.NewDisplay(bus)
	wd.RegisterObserver(display, 0)

	statsDisplay := observers.NewStatsDisplay(bus)
	wd.RegisterObserver(statsDisplay, 1)

	wd.SetMeasurements(3, 0.7, 760)
	wd.SetMeasurements(4, 0.8, 761)
	wd.RemoveObserver(statsDisplay)
	wd.SetMeasurements(10, 0.8, 761)
	wd.SetMeasurements(-10, 0.8, 761)
}
