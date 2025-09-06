package duck

import (
	"fmt"

	"simuduck/pkg/behavior"
)

func NewRedheadDuck() Duck {
	return &redheadDuck{
		duck: duck{
			flyBehavior:   behavior.NewFlyWithWingsBehavior(),
			quackBehavior: behavior.NewQuackBehavior(),
			danceBehavior: behavior.NewDanceMinuetBehavior(),
		},
	}
}

type redheadDuck struct {
	duck
}

func (d *redheadDuck) Display() {
	fmt.Println("I'm readhead duck")
}
