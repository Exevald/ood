package model

type WeatherInfo struct {
	Temperature float64
	Humidity    float64
	Pressure    float64
}

func NewWeatherData() *WeatherData {
	return &WeatherData{
		Observable: NewObservable(),
		pressure:   760,
	}
}

type WeatherData struct {
	Observable

	temperature float64
	humidity    float64
	pressure    float64
}

func (wd *WeatherData) GetTemperature() float64 {
	return wd.temperature
}

func (wd *WeatherData) GetHumidity() float64 {
	return wd.humidity
}

func (wd *WeatherData) GetPressure() float64 {
	return wd.pressure
}

func (wd *WeatherData) MeasurementsChanged() {
	data := WeatherInfo{
		Temperature: wd.GetTemperature(),
		Humidity:    wd.GetHumidity(),
		Pressure:    wd.GetPressure(),
	}
	wd.NotifyObservers(data)
}

func (wd *WeatherData) SetMeasurements(temp, humidity, pressure float64) {
	wd.temperature = temp
	wd.humidity = humidity
	wd.pressure = pressure
	wd.MeasurementsChanged()
}
