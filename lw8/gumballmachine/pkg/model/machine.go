package model

import (
	"fmt"

	"slides/pkg/model/states"
)

type GumballMachine interface {
	InsertQuarter()
	EjectQuarter()
	TurnCrank()
	Refill(count int)

	GetBallCount() int
	GetQuarterCount() int
	String() string
}

func NewGumballMachine(numberGumballs int) GumballMachine {
	machine := &gumballMachine{
		count: numberGumballs,
	}

	machine.soldOutState = states.NewSoldOutState(machine)
	machine.noQuarterState = states.NewNoQuarterState(machine)
	machine.hasQuarterState = states.NewHasQuarterState(machine)
	machine.soldState = states.NewSoldState(machine)

	if numberGumballs > 0 {
		machine.currentState = machine.noQuarterState
	} else {
		machine.currentState = machine.soldOutState
	}

	return machine
}

type gumballMachine struct {
	soldOutState    states.State
	noQuarterState  states.State
	hasQuarterState states.State
	soldState       states.State

	currentState states.State
	count        int
	quarterCount int
}

func (g *gumballMachine) InsertQuarter() {
	g.currentState.InsertQuarter()
}

func (g *gumballMachine) EjectQuarter() {
	g.currentState.EjectQuarter()
}

func (g *gumballMachine) TurnCrank() {
	g.currentState.TurnCrank()
	g.currentState.Dispense()
}

func (g *gumballMachine) Refill(count int) {
	g.currentState.Refill(count)
}

func (g *gumballMachine) String() string {
	return fmt.Sprintf(
		"\nMighty Gumball, Inc.\nGo-enabled Standing Gumball Model #2025\nInventory: %d gumballs, %d quarters\nMachine is %s\n",
		g.count, g.quarterCount, g.currentState.String(),
	)
}

func (g *gumballMachine) SetState(state states.State) {
	g.currentState = state
}

func (g *gumballMachine) ReleaseBall() {
	if g.count > 0 {
		fmt.Println("A gumball comes rolling out the slot...")
		g.count--
	}
}

func (g *gumballMachine) AddQuarter() {
	g.quarterCount++
}

func (g *gumballMachine) UseQuarter() {
	if g.quarterCount > 0 {
		g.quarterCount--
	}
}

func (g *gumballMachine) ReturnQuarters() {
	fmt.Printf("Returning %d coin(s)\n", g.quarterCount)
	g.quarterCount = 0
}

func (g *gumballMachine) AddBalls(count int) {
	g.count += count
	fmt.Printf("Refilled. New ball count: %d\n", g.count)
}

func (g *gumballMachine) GetBallCount() int {
	return g.count
}

func (g *gumballMachine) GetQuarterCount() int {
	return g.quarterCount
}

func (g *gumballMachine) GetSoldOutState() states.State {
	return g.soldOutState
}

func (g *gumballMachine) GetNoQuarterState() states.State {
	return g.noQuarterState
}

func (g *gumballMachine) GetHasQuarterState() states.State {
	return g.hasQuarterState
}

func (g *gumballMachine) GetSoldState() states.State {
	return g.soldState
}
