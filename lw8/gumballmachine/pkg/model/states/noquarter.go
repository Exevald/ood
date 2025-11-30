package states

import "fmt"

func NewNoQuarterState(machine stateContext) State {
	return &noQuarterState{machine: machine}
}

type noQuarterState struct {
	machine stateContext
}

func (s *noQuarterState) InsertQuarter() {
	fmt.Println("You inserted a quarter")
	s.machine.AddQuarter()
	s.machine.SetState(s.machine.GetHasQuarterState())
}

func (s *noQuarterState) EjectQuarter() {
	fmt.Println("You haven't inserted a quarter")
}

func (s *noQuarterState) TurnCrank() {
	fmt.Println("You turned but there's no quarter")
}

func (s *noQuarterState) Dispense() {
	fmt.Println("You need to pay first")
}

func (s *noQuarterState) Refill(count int) {
	s.machine.AddBalls(count)
}

func (s *noQuarterState) String() string {
	return "waiting for quarter"
}
