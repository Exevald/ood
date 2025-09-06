package main

import (
	"fmt"
	"simuduck/pkg/behavior"

	"simuduck/pkg/duck"
)

func DrawDuck(duck duck.Duck) {
	duck.Display()
}

func PlayWithDuck(duck duck.Duck) {
	DrawDuck(duck)
	duck.Quack()
	duck.Fly()
	duck.Fly()
	duck.Dance()

	fmt.Printf("\n")
}

func main() {
	mallardDuck := duck.NewMallardDuck()
	PlayWithDuck(mallardDuck)

	redheadDuck := duck.NewRedheadDuck()
	PlayWithDuck(redheadDuck)

	rubberDuck := duck.NewRubberDuck()
	PlayWithDuck(rubberDuck)

	decoyDuck := duck.NewDecoyDuck()
	PlayWithDuck(decoyDuck)

	modelDuck := duck.NewModelDuck()
	PlayWithDuck(modelDuck)

	modelDuck.SetFlyBehavior(behavior.NewFlyWithWingsBehavior())
	PlayWithDuck(modelDuck)
}
