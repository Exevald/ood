package model

import (
	"sort"
	"sync"
)

type Observer interface {
	Update(subjectID string, data WeatherInfo)
}

type Observable interface {
	RegisterObserver(observer Observer, priority int)
	RemoveObserver(observer Observer)
	NotifyObservers(data WeatherInfo)
}

type observerEntry struct {
	observer Observer
	priority int
}

type observable struct {
	mu         sync.RWMutex
	observers  map[Observer]int
	sortedList []observerEntry
}

func NewObservable() Observable {
	return &observable{
		observers: make(map[Observer]int),
	}
}

func (o *observable) RegisterObserver(observer Observer, priority int) {
	o.mu.Lock()
	defer o.mu.Unlock()
	if _, exists := o.observers[observer]; exists {
		return
	}
	o.observers[observer] = priority
	o.rebuildSortedList()
}

func (o *observable) RemoveObserver(observer Observer) {
	o.mu.Lock()
	defer o.mu.Unlock()
	if _, exists := o.observers[observer]; !exists {
		return
	}
	delete(o.observers, observer)
	o.rebuildSortedList()
}

func (o *observable) rebuildSortedList() {
	o.sortedList = make([]observerEntry, 0, len(o.observers))
	for obs, prio := range o.observers {
		o.sortedList = append(o.sortedList, observerEntry{observer: obs, priority: prio})
	}
	sort.Slice(o.sortedList, func(i, j int) bool {
		return o.sortedList[i].priority > o.sortedList[j].priority
	})
}

func (o *observable) NotifyObservers(data WeatherInfo) {
	o.mu.RLock()
	listCopy := make([]observerEntry, len(o.sortedList))
	copy(listCopy, o.sortedList)
	o.mu.RUnlock()

	for _, entry := range listCopy {
		entry.observer.Update("", data)
	}
}
