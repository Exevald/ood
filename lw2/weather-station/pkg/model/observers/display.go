package observers

import (
	"fmt"

	"weatherstation/pkg/model"
)

type display struct{}

func NewDisplay() model.Observer {
	return &display{}
}

func (d *display) Update(subjectID string, data model.WeatherInfo) {
	fmt.Printf("Subject %s Current Temp %.2f\n", subjectID, data.Temperature)
	fmt.Printf("Subject %s Current Hum %.2f\n", subjectID, data.Humidity)
	fmt.Printf("Subject %s Current Pressure %.2f\n", subjectID, data.Pressure)
	fmt.Println("-----------------")
}
