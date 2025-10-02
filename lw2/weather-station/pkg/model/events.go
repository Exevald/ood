package model

type EventType string

const (
	EventTemperatureChanged EventType = "temperature_changed"
	EventHumidityChanged    EventType = "humidity_changed"
	EventPressureChanged    EventType = "pressure_changed"
	EventWindSpeedChanged   EventType = "windspeed_changed"
	EventWindDirChanged     EventType = "winddir_changed"
)
