package model

import (
	"sort"
	"sync"
	"sync/atomic"
)

type EventHandler func(event Event)

type Event struct {
	Type EventType
	Data interface{}
}

type subscriptionID int64

var globalID int64

type eventEntry struct {
	id       subscriptionID
	handler  EventHandler
	priority int
}

type EventBus interface {
	Subscribe(eventType EventType, handler EventHandler, priority int) subscriptionID
	Unsubscribe(eventType EventType, id subscriptionID)
	Dispatch(event Event)
}

type eventBus struct {
	mu            sync.RWMutex
	eventEntryMap map[EventType][]eventEntry
}

func NewEventBus() EventBus {
	return &eventBus{
		eventEntryMap: make(map[EventType][]eventEntry),
	}
}

func (b *eventBus) Subscribe(eventType EventType, handler EventHandler, priority int) subscriptionID {
	id := b.nextID()
	b.mu.Lock()
	defer b.mu.Unlock()

	entry := eventEntry{id: id, handler: handler, priority: priority}
	b.eventEntryMap[eventType] = append(b.eventEntryMap[eventType], entry)
	b.sortHandlers(eventType)
	return id
}

func (b *eventBus) Unsubscribe(eventType EventType, id subscriptionID) {
	b.mu.Lock()
	defer b.mu.Unlock()

	handlers := b.eventEntryMap[eventType]
	for i, entry := range handlers {
		if entry.id == id {
			b.eventEntryMap[eventType] = append(handlers[:i], handlers[i+1:]...)
			return
		}
	}
}

func (b *eventBus) Dispatch(event Event) {
	b.mu.RLock()
	handlersOnThisType := make([]eventEntry, len(b.eventEntryMap[event.Type]))
	copy(handlersOnThisType, b.eventEntryMap[event.Type])
	b.mu.RUnlock()

	for _, entry := range handlersOnThisType {
		entry.handler(event)
	}
}

func (b *eventBus) sortHandlers(eventType EventType) {
	handlers := b.eventEntryMap[eventType]
	sort.Slice(handlers, func(i, j int) bool {
		return handlers[i].priority > handlers[j].priority
	})
	b.eventEntryMap[eventType] = handlers
}

func (b *eventBus) nextID() subscriptionID {
	return subscriptionID(atomic.AddInt64(&globalID, 1))
}
