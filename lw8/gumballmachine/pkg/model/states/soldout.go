package states

import (
	"fmt"
)

func NewSoldOutState(machine stateContext) State {
	return &soldOutState{machine: machine}
}

type soldOutState struct {
	machine stateContext
}

func (s *soldOutState) InsertQuarter() {
	fmt.Println("You can't insert a quarter, the machine is sold out")
}

func (s *soldOutState) EjectQuarter() {
	if s.machine.GetQuarterCount() > 0 {
		s.machine.ReturnQuarters()
		s.machine.SetState(s.machine.GetNoQuarterState())
	} else {
		fmt.Println("You can't eject, you haven't inserted a quarter yet")
	}
}

func (s *soldOutState) TurnCrank() {
	fmt.Println("You turned but there's no gumballs")
}

func (s *soldOutState) Dispense() {
	fmt.Println("No gumball dispensed")
}

func (s *soldOutState) Refill(count int) {
	s.machine.AddBalls(count)
	if s.machine.GetBallCount() > 0 {
		if s.machine.GetQuarterCount() > 0 {
			s.machine.SetState(s.machine.GetHasQuarterState())
		} else {
			s.machine.SetState(s.machine.GetNoQuarterState())
		}
	}
}

func (s *soldOutState) String() string {
	return "sold out"
}
