package duck

import (
	"fmt"

	"simuduck/pkg/behavior"
)

func NewRubberDuck() Duck {
	return &rubberDuck{
		duck: duck{
			flyBehavior:   behavior.NewFlyNoWayBehavior(),
			quackBehavior: behavior.NewSqueakBehavior(),
			danceBehavior: behavior.NewNoDanceBehavior(),
		},
	}
}

type rubberDuck struct {
	duck
}

func (d *rubberDuck) Display() {
	fmt.Println("I'm rubber duck")
}
