package model

type WeatherInfo struct {
	Temperature float64
	Humidity    float64
	Pressure    float64
}

type BaseWeatherData struct {
	Observable

	name        string
	temperature float64
	humidity    float64
	pressure    float64
	bus         EventBus
}

func (b *BaseWeatherData) GetTemperature() float64 {
	return b.temperature
}

func (b *BaseWeatherData) GetHumidity() float64 {
	return b.humidity
}

func (b *BaseWeatherData) GetPressure() float64 {
	return b.pressure
}

type IndoorWeatherData struct {
	BaseWeatherData
}

func NewIndoorWeatherData(name string, eventBus EventBus) *IndoorWeatherData {
	return &IndoorWeatherData{
		BaseWeatherData: BaseWeatherData{
			Observable: NewObservable(),
			name:       name,
			bus:        eventBus,
		},
	}
}

func (wd *IndoorWeatherData) SetMeasurements(temp, humidity, pressure float64) {
	wd.temperature = temp
	wd.humidity = humidity
	wd.pressure = pressure
	wd.NotifyObservers(wd.name, WeatherInfo{
		Temperature: temp,
		Humidity:    humidity,
		Pressure:    pressure,
	})

	wd.bus.Dispatch(Event{Type: EventTemperatureChanged, Data: temp})
	wd.bus.Dispatch(Event{Type: EventHumidityChanged, Data: humidity})
	wd.bus.Dispatch(Event{Type: EventPressureChanged, Data: pressure})
}

type OutdoorWeatherData struct {
	BaseWeatherData

	windSpeed float64
	windDir   string
}

func NewOutdoorWeatherData(name string) *OutdoorWeatherData {
	return &OutdoorWeatherData{
		BaseWeatherData: BaseWeatherData{
			Observable: NewObservable(),
			name:       name,
		},
	}
}

func (wd *OutdoorWeatherData) GetWindSpeed() float64 {
	return wd.windSpeed
}

func (wd *OutdoorWeatherData) GetWindDir() string {
	return wd.windDir
}

func (wd *OutdoorWeatherData) SetMeasurements(temp, humidity, pressure, windSpeed float64, windDir string) {
	wd.temperature = temp
	wd.humidity = humidity
	wd.pressure = pressure
	wd.windSpeed = windSpeed
	wd.windDir = windDir

	wd.bus.Dispatch(Event{Type: EventTemperatureChanged, Data: temp})
	wd.bus.Dispatch(Event{Type: EventHumidityChanged, Data: humidity})
	wd.bus.Dispatch(Event{Type: EventPressureChanged, Data: pressure})
	wd.bus.Dispatch(Event{Type: EventWindSpeedChanged, Data: windSpeed})
	wd.bus.Dispatch(Event{Type: EventWindDirChanged, Data: windDir})
}
