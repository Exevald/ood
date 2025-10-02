package tests

import (
	"testing"

	"weatherstation/pkg/model"
)

func TestRegisterRemoveNotifyObservers(t *testing.T) {
	o := model.NewObservable()

	var calls []string

	o1 := &mockObserver{id: "obs1", called: &calls}
	o2 := &mockObserver{id: "obs2", called: &calls}
	o3 := &mockObserver{id: "obs3", called: &calls}

	o.RegisterObserver(o1, 5)
	o.RegisterObserver(o2, 10)
	o.RegisterObserver(o3, 1)

	o.RegisterObserver(o1, 15)

	o.NotifyObservers("test", model.WeatherInfo{Temperature: 0})

	expectedOrder := []string{"obs2", "obs1", "obs3"}
	for i, v := range expectedOrder {
		if calls[i] != v {
			t.Errorf("Notify order incorrect at %d: got %s expected %s", i, calls[i], v)
		}
	}

	o.RemoveObserver(o1)
	calls = calls[:0]

	o.NotifyObservers("test", model.WeatherInfo{Temperature: 0})

	expectedOrder = []string{"obs2", "obs3"}
	if len(calls) != len(expectedOrder) {
		t.Fatalf("Expected %d calls after removal but got %d", len(expectedOrder), len(calls))
	}
	for i, v := range expectedOrder {
		if calls[i] != v {
			t.Errorf("Notify order incorrect at %d after removal: got %s expected %s", i, calls[i], v)
		}
	}
}

func TestSafeRemoveSelfInUpdate(t *testing.T) {
	o := model.NewObservable()

	var calls []string

	var obs1 *mockObserver
	obs1 = &mockObserver{
		id: "obs1", called: &calls,
		rm: func() {
			o.RemoveObserver(obs1)
		},
	}
	obs2 := &mockObserver{id: "obs2", called: &calls}

	o.RegisterObserver(obs1, 1)
	o.RegisterObserver(obs2, 2)

	o.NotifyObservers("test", model.WeatherInfo{Temperature: 0})

	expectedCalls := []string{"obs2", "obs1"}
	if len(calls) != len(expectedCalls) {
		t.Fatalf("Expected %d calls but got %d", len(expectedCalls), len(calls))
	}
	for i, v := range expectedCalls {
		if calls[i] != v {
			t.Errorf("Call order mismatch at %d: got %s expected %s", i, calls[i], v)
		}
	}

	calls = calls[:0]
	o.NotifyObservers("test", model.WeatherInfo{Temperature: 0})
	if len(calls) != 1 || calls[0] != "obs2" {
		t.Errorf("After removal, expected only 'obs2' call, got %v", calls)
	}
}

func TestRemoveNonexistentObserver(t *testing.T) {
	o := model.NewObservable()
	obs := &mockObserver{id: "obs"}

	o.RemoveObserver(obs)

	o.RegisterObserver(obs, 0)
	o.RemoveObserver(obs)
	o.RemoveObserver(obs)
}
