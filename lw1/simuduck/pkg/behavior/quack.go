package behavior

import "fmt"

type QuackBehavior interface {
	Quack()
}

func NewQuackBehavior() QuackBehavior {
	return &quackBehavior{}
}

type quackBehavior struct{}

func (q *quackBehavior) Quack() {
	fmt.Println("Quack")
}

func NewSqueakBehavior() QuackBehavior {
	return &squeakBehavior{}
}

type squeakBehavior struct{}

func (s *squeakBehavior) Quack() {
	fmt.Println("Squeak")
}

func NewMuteQuackBehavior() QuackBehavior {
	return &muteQuackBehavior{}
}

type muteQuackBehavior struct{}

func (m *muteQuackBehavior) Quack() {
}
