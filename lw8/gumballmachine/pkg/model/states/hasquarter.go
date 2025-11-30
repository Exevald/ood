package states

import "fmt"

func NewHasQuarterState(machine stateContext) State {
	return &hasQuarterState{machine: machine}
}

type hasQuarterState struct {
	machine stateContext
}

func (s *hasQuarterState) InsertQuarter() {
	if s.machine.GetQuarterCount() < 5 {
		fmt.Println("You inserted another quarter")
		s.machine.AddQuarter()
	} else {
		fmt.Println("You can't insert another quarter, the slot is full")
	}
}

func (s *hasQuarterState) EjectQuarter() {
	s.machine.ReturnQuarters()
	s.machine.SetState(s.machine.GetNoQuarterState())
}

func (s *hasQuarterState) TurnCrank() {
	fmt.Println("You turned...")
	s.machine.UseQuarter()
	s.machine.SetState(s.machine.GetSoldState())
}

func (s *hasQuarterState) Dispense() {
	fmt.Println("No gumball dispensed")
}

func (s *hasQuarterState) Refill(count int) {
	s.machine.AddBalls(count)
}

func (s *hasQuarterState) String() string {
	return "waiting for turn of crank"
}
