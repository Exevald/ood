package tests

import (
	"sync"

	"weatherstation/pkg/model"
)

type mockObserver struct {
	id     string
	rm     func()
	lock   sync.Mutex
	called *[]string
}

func (o *mockObserver) Update(_ string, _ model.WeatherInfo) {
	o.lock.Lock()
	defer o.lock.Unlock()
	*o.called = append(*o.called, o.id)
	if o.rm != nil {
		o.rm()
		o.rm = nil
	}
}
