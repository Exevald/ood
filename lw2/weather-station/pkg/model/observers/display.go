package observers

import (
	"fmt"

	"weatherstation/pkg/model"
)

func NewDisplay(bus model.EventBus) model.Observer {
	display := &display{}

	bus.Subscribe(model.EventTemperatureChanged, display.onTemperatureChanged, 0)
	bus.Subscribe(model.EventPressureChanged, display.onPressureChanged, 1)

	return display
}

type display struct{}

func (d *display) onTemperatureChanged(event model.Event) {
	fmt.Printf("Current Temperature %.2f\n", event.Data)
}

func (d *display) onPressureChanged(event model.Event) {
	fmt.Printf("Current Pressure %.2f\n", event.Data)
}
