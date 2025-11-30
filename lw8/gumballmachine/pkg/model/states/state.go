package states

type State interface {
	InsertQuarter()
	EjectQuarter()
	TurnCrank()
	Dispense()
	Refill(count int)
	String() string
}

type stateContext interface {
	SetState(State)

	GetSoldOutState() State
	GetNoQuarterState() State
	GetHasQuarterState() State
	GetSoldState() State

	GetBallCount() int
	GetQuarterCount() int

	AddQuarter()
	UseQuarter()
	ReturnQuarters()
	ReleaseBall()
	AddBalls(count int)
}
