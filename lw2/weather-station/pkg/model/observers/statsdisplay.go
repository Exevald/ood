package observers

import (
	"fmt"

	"weatherstation/pkg/model"
)

func NewStatsDisplay(bus model.EventBus) model.Observer {
	display := &statsDisplay{
		temperature: model.NewStats(),
		humidity:    model.NewStats(),
		pressure:    model.NewStats(),
	}
	bus.Subscribe(model.EventTemperatureChanged, display.onTemperatureChanged, 0)
	bus.Subscribe(model.EventHumidityChanged, display.onHumidityChanged, 1)
	bus.Subscribe(model.EventPressureChanged, display.onPressureChanged, 2)

	return display
}

type statsDisplay struct {
	temperature model.Stats
	humidity    model.Stats
	pressure    model.Stats
}

func (d *statsDisplay) onTemperatureChanged(event model.Event) {
	d.temperature.Update(event.Data.(float64))
	fmt.Println(d.temperature.ToString("Temperature"))
}

func (d *statsDisplay) onHumidityChanged(event model.Event) {
	d.humidity.Update(event.Data.(float64))
	fmt.Println(d.humidity.ToString("Humidity"))
}

func (d *statsDisplay) onPressureChanged(event model.Event) {
	d.pressure.Update(event.Data.(float64))
	fmt.Println(d.pressure.ToString("Pressure"))
}
