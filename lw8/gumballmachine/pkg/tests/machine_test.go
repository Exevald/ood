package tests

import (
	"strings"
	"testing"

	"slides/pkg/model"
)

func TestInitialState(t *testing.T) {
	gm := model.NewGumballMachine(10)

	if gm.GetBallCount() != 10 {
		t.Errorf("Expected 10 balls, got %d", gm.GetBallCount())
	}
	assertState(t, gm, "waiting for quarter")
}

func TestSoldOutStart(t *testing.T) {
	gm := model.NewGumballMachine(0)
	assertState(t, gm, "sold out")
}

func TestStandardCycle(t *testing.T) {
	gm := model.NewGumballMachine(5)

	gm.InsertQuarter()
	if gm.GetQuarterCount() != 1 {
		t.Errorf("Expected 1 quarter, got %d", gm.GetQuarterCount())
	}
	assertState(t, gm, "waiting for turn of crank")

	gm.TurnCrank()

	if gm.GetBallCount() != 4 {
		t.Errorf("Expected 4 balls, got %d", gm.GetBallCount())
	}
	if gm.GetQuarterCount() != 0 {
		t.Errorf("Expected 0 quarters, got %d", gm.GetQuarterCount())
	}
	assertState(t, gm, "waiting for quarter")
}

func TestMultiQuarterLimit(t *testing.T) {
	gm := model.NewGumballMachine(10)

	for i := 0; i < 5; i++ {
		gm.InsertQuarter()
	}

	if gm.GetQuarterCount() != 5 {
		t.Errorf("Expected 5 quarters, got %d", gm.GetQuarterCount())
	}

	gm.InsertQuarter()
	if gm.GetQuarterCount() != 5 {
		t.Errorf("Limit should be 5, got %d", gm.GetQuarterCount())
	}
}

func TestEjectReturnsAllQuarters(t *testing.T) {
	gm := model.NewGumballMachine(5)
	gm.InsertQuarter()
	gm.InsertQuarter()
	gm.InsertQuarter()

	gm.EjectQuarter()

	if gm.GetQuarterCount() != 0 {
		t.Errorf("Expected 0 quarters after eject, got %d", gm.GetQuarterCount())
	}
	assertState(t, gm, "waiting for quarter")
}

func TestTurnCrankWithRemainingCoins(t *testing.T) {
	gm := model.NewGumballMachine(5)

	gm.InsertQuarter()
	gm.InsertQuarter()

	gm.TurnCrank()

	if gm.GetBallCount() != 4 {
		t.Errorf("Expected 4 balls left")
	}
	if gm.GetQuarterCount() != 1 {
		t.Errorf("Expected 1 quarter left")
	}
	assertState(t, gm, "waiting for turn of crank")

	gm.TurnCrank()

	if gm.GetQuarterCount() != 0 {
		t.Errorf("Expected 0 quarters left")
	}
	assertState(t, gm, "waiting for quarter")
}

func TestCoinsMoreThanBalls(t *testing.T) {
	gm := model.NewGumballMachine(1)
	gm.InsertQuarter()
	gm.InsertQuarter()

	gm.TurnCrank()

	if gm.GetBallCount() != 0 {
		t.Errorf("Balls should be 0")
	}
	if gm.GetQuarterCount() != 1 {
		t.Errorf("Should have 1 coin credit left")
	}
	assertState(t, gm, "sold out")

	gm.EjectQuarter()

	if gm.GetQuarterCount() != 0 {
		t.Errorf("Coin should be returned")
	}

	assertState(t, gm, "waiting for quarter")
}

func TestRefill(t *testing.T) {
	gm := model.NewGumballMachine(0)
	assertState(t, gm, "sold out")

	gm.Refill(5)

	if gm.GetBallCount() != 5 {
		t.Errorf("Refill failed")
	}
	assertState(t, gm, "waiting for quarter")
}

func TestRefillWithCoins(t *testing.T) {
	gm := model.NewGumballMachine(1)
	gm.InsertQuarter()

	gm.Refill(10)

	if gm.GetBallCount() != 11 {
		t.Errorf("Refill should add up")
	}
	if gm.GetQuarterCount() != 1 {
		t.Errorf("Coins should be preserved")
	}
	assertState(t, gm, "waiting for turn of crank")
}

func assertState(t *testing.T, gm model.GumballMachine, expectedText string) {
	t.Helper()
	if !strings.Contains(gm.String(), expectedText) {
		t.Errorf("Expected state to contain '%s', but got:\n%s", expectedText, gm.String())
	}
}
