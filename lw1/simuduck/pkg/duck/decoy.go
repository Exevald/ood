package duck

import (
	"fmt"

	"simuduck/pkg/behavior"
)

func NewDecoyDuck() Duck {
	return &decoyDuck{
		duck: duck{
			flyBehavior:   behavior.NewFlyNoWayBehavior(),
			quackBehavior: behavior.NewMuteQuackBehavior(),
			danceBehavior: behavior.NewNoDanceBehavior(),
		},
	}
}

type decoyDuck struct {
	duck
}

func (d *decoyDuck) Display() {
	fmt.Println("I'm decoy duck")
}
