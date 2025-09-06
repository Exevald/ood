package behavior

import "fmt"

type DanceBehavior interface {
	Dance()
}

func NewDanceWaltzBehavior() DanceBehavior {
	return &danceWaltzBehavior{}
}

type danceWaltzBehavior struct{}

func (d *danceWaltzBehavior) Dance() {
	fmt.Println("Dancing waltz")
}

func NewDanceMinuetBehavior() DanceBehavior {
	return &danceMinuetBehavior{}
}

type danceMinuetBehavior struct{}

func (d *danceMinuetBehavior) Dance() {
	fmt.Println("Dancing minuet")
}

func NewNoDanceBehavior() DanceBehavior {
	return &noDanceBehavior{}
}

type noDanceBehavior struct{}

func (n *noDanceBehavior) Dance() {
}
