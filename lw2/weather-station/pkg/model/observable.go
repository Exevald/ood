package model

type Observer interface {
	Update(data WeatherInfo)
}

type Observable interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers(data WeatherInfo)
}

type observable struct {
	observers map[Observer]struct{}
}

func NewObservable() Observable {
	return &observable{observers: make(map[Observer]struct{})}
}

func (o *observable) RegisterObserver(observer Observer) {
	o.observers[observer] = struct{}{}
}

func (o *observable) RemoveObserver(observer Observer) {
	delete(o.observers, observer)
}

func (o *observable) NotifyObservers(data WeatherInfo) {
	for observer := range o.observers {
		observer.Update(data)
	}
}
