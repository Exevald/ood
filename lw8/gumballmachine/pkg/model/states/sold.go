package states

import "fmt"

func NewSoldState(machine stateContext) State {
	return &soldState{machine: machine}
}

type soldState struct {
	machine stateContext
}

func (s *soldState) InsertQuarter() {
	fmt.Println("Please wait, we're already giving you a gumball")
}

func (s *soldState) EjectQuarter() {
	fmt.Println("Sorry you already turned the crank")
}

func (s *soldState) TurnCrank() {
	fmt.Println("Turning twice doesn't get you another gumball")
}

func (s *soldState) Dispense() {
	s.machine.ReleaseBall()

	if s.machine.GetBallCount() == 0 {
		fmt.Println("Oops, out of gumballs")
		s.machine.SetState(s.machine.GetSoldOutState())
	} else {
		if s.machine.GetQuarterCount() > 0 {
			s.machine.SetState(s.machine.GetHasQuarterState())
		} else {
			s.machine.SetState(s.machine.GetNoQuarterState())
		}
	}
}
func (s *soldState) Refill(_ int) {
	fmt.Println("Cannot refill while dispensing a gumball")
}

func (s *soldState) String() string {
	return "delivering a gumball"
}
