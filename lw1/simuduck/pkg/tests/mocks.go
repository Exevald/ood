package tests

type MockFlyWithWingsBehavior struct {
	flightCount int
}

func (m *MockFlyWithWingsBehavior) Fly() {
	m.flightCount++
}

func (m *MockFlyWithWingsBehavior) GetFlightCount() int {
	return m.flightCount
}

type MockFlyNoWayBehavior struct {
}

func (m *MockFlyNoWayBehavior) Fly() {
}

func (m *MockFlyNoWayBehavior) GetFlightCount() int {
	return 0
}

type MockDanceWaltzBehavior struct {
	danceCount int
}

func (m *MockDanceWaltzBehavior) Dance() {
	m.danceCount++
}

func (m *MockDanceWaltzBehavior) GetDanceCount() int {
	return m.danceCount
}

type MockDanceMinuetBehavior struct {
	danceCount int
}

func (m *MockDanceMinuetBehavior) Dance() {
	m.danceCount++
}

func (m *MockDanceMinuetBehavior) GetDanceCount() int {
	return m.danceCount
}

type MockNoDanceBehavior struct {
}

func (m *MockNoDanceBehavior) Dance() {
}

func (m *MockNoDanceBehavior) GetDanceCount() int {
	return 0
}

type MockQuackBehavior struct {
	quackCount int
}

func (m *MockQuackBehavior) Quack() {
	m.quackCount++
}

func (m *MockQuackBehavior) GetQuackCount() int {
	return m.quackCount
}

type MockSqueakBehavior struct {
	quackCount int
}

func (m *MockSqueakBehavior) Quack() {
	m.quackCount++
}

func (m *MockSqueakBehavior) GetQuackCount() int {
	return m.quackCount
}

type MockMuteQuackBehavior struct {
}

func (m *MockMuteQuackBehavior) Quack() {
}

func (m *MockMuteQuackBehavior) GetQuackCount() int {
	return 0
}
