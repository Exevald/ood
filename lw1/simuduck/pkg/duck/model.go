package duck

import (
	"fmt"

	"simuduck/pkg/behavior"
)

func NewModelDuck() Duck {
	return &modelDuck{
		duck: duck{
			flyBehavior:   behavior.NewFlyNoWayBehavior(),
			quackBehavior: behavior.NewQuackBehavior(),
			danceBehavior: behavior.NewNoDanceBehavior(),
		},
	}
}

type modelDuck struct {
	duck
}

func (d *modelDuck) Display() {
	fmt.Println("I'm model duck")
}
