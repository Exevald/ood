package tests

import (
	"testing"
)

func TestFlyBehavior(t *testing.T) {
	flyWithWingsBehavior := MockFlyWithWingsBehavior{}
	if flyWithWingsBehavior.GetFlightCount() != 0 {
		t.Error("The initial number of flights should be 0")
	}

	flyWithWingsBehavior.Fly()
	if flyWithWingsBehavior.GetFlightCount() != 1 {
		t.Error("The flight count after flight should be 1")
	}

	flyNoWayBehavior := MockFlyNoWayBehavior{}
	if flyNoWayBehavior.GetFlightCount() != 0 {
		t.Error("The initial number of flights should be 0")
	}

	flyNoWayBehavior.Fly()
	if flyNoWayBehavior.GetFlightCount() != 0 {
		t.Error("The flight count after no way flight should be 0")
	}
}

func TestQuackBehavior(t *testing.T) {
	quackBehavior := MockQuackBehavior{}
	if quackBehavior.GetQuackCount() != 0 {
		t.Error("The initial quack count should be 0")
	}
	quackBehavior.Quack()
	if quackBehavior.GetQuackCount() != 1 {
		t.Error("The quack count after quack should be 1")
	}

	squeakBehavior := MockSqueakBehavior{}
	if squeakBehavior.GetQuackCount() != 0 {
		t.Error("The initial squeak count should be 0")
	}
	squeakBehavior.Quack()
	if squeakBehavior.GetQuackCount() != 1 {
		t.Error("The quack count after quack should be 1")
	}

	muteQuackBehavior := MockMuteQuackBehavior{}
	if muteQuackBehavior.GetQuackCount() != 0 {
		t.Error("The initial mute quack count should be 0")
	}
	muteQuackBehavior.Quack()
	if muteQuackBehavior.GetQuackCount() != 0 {
		t.Error("The quack count after mute quack should be 0")
	}
}

func TestDanceBehavior(t *testing.T) {
	danceWaltzBehavior := MockDanceWaltzBehavior{}
	if danceWaltzBehavior.GetDanceCount() != 0 {
		t.Error("The initial dance count should be 0")
	}
	danceWaltzBehavior.Dance()
	if danceWaltzBehavior.GetDanceCount() != 1 {
		t.Error("The dance count after dance should be 1")
	}

	danceMinuetBehavior := MockDanceMinuetBehavior{}
	if danceMinuetBehavior.GetDanceCount() != 0 {
		t.Error("The initial dance count should be 0")
	}
	danceMinuetBehavior.Dance()
	if danceMinuetBehavior.GetDanceCount() != 1 {
		t.Error("The dance count after dance should be 1")
	}

	noDanceBehavior := MockNoDanceBehavior{}
	if noDanceBehavior.GetDanceCount() != 0 {
		t.Error("The initial dance count should be 0")
	}
	noDanceBehavior.Dance()
	if noDanceBehavior.GetDanceCount() != 0 {
		t.Error("The dance count after no dance should be 0")
	}
}
