package duck

import (
	"fmt"

	"simuduck/pkg/behavior"
)

type Duck interface {
	Quack()
	Fly()
	Dance()
	Swim()
	Display()
	SetFlyBehavior(behavior behavior.FlyBehavior)
}

type duck struct {
	flyBehavior   behavior.FlyBehavior
	quackBehavior behavior.QuackBehavior
	danceBehavior behavior.DanceBehavior
}

func (d *duck) Quack() {
	if d.quackBehavior != nil {
		d.quackBehavior.Quack()
	}
}

func (d *duck) Fly() {
	if d.flyBehavior == nil {
		return
	}

	d.flyBehavior.Fly()

	if d.flyBehavior.GetFlightCount()%2 == 0 {
		d.quackBehavior.Quack()
	}
}

func (d *duck) Dance() {
	if d.danceBehavior != nil {
		d.danceBehavior.Dance()
	}
}

func (d *duck) Swim() {
	fmt.Println("I'm swimming")
}

func (d *duck) Display() {
	fmt.Println("I'm base duck")
}

func (d *duck) SetFlyBehavior(behavior behavior.FlyBehavior) {
	if behavior != nil {
		d.flyBehavior = behavior
	}
}
