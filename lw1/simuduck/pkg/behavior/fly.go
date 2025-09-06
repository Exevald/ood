package behavior

import "fmt"

type FlyBehavior interface {
	Fly()
	GetFlightCount() int
}

func NewFlyWithWingsBehavior() FlyBehavior {
	return &flyWithWingsBehavior{}
}

type flyWithWingsBehavior struct {
	flightCount int
}

func (f *flyWithWingsBehavior) Fly() {
	f.flightCount++
	fmt.Println("Flight with number: ", f.flightCount)
}

func (f *flyWithWingsBehavior) GetFlightCount() int {
	return f.flightCount
}

func NewFlyNoWayBehavior() FlyBehavior {
	return &flyNoWayBehavior{}
}

type flyNoWayBehavior struct{}

func (f *flyNoWayBehavior) Fly() {
}

func (f *flyNoWayBehavior) GetFlightCount() int {
	return 0
}
