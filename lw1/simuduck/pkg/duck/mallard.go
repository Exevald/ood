package duck

import (
	"fmt"

	"simuduck/pkg/behavior"
)

func NewMallardDuck() Duck {
	return &mallardDuck{
		duck: duck{
			flyBehavior:   behavior.NewFlyWithWingsBehavior(),
			quackBehavior: behavior.NewQuackBehavior(),
			danceBehavior: behavior.NewDanceWaltzBehavior(),
		},
	}
}

type mallardDuck struct {
	duck
}

func (d *mallardDuck) Display() {
	fmt.Println("I'm mallard duck")
}
