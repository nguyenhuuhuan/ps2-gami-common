package events

import (
	"github.com/asaskevich/EventBus"
)

var eventManager EventBus.Bus

// Init Event Manager with event bus
func Init() *EventBus.Bus {
	eventManager = EventBus.New()
	return &eventManager
}

// AddSync Event Handler to Event Bus
func AddSync(topic string, fn interface{}) {
	_ = eventManager.Subscribe(topic, fn)
}

// AddAsync Handler to Event Bus
func AddAsync(topic string, fn interface{}) {
	_ = eventManager.SubscribeAsync(topic, fn, false)
}

// Publish an event to event bus
func Publish(topic string, args ...interface{}) {
	eventManager.Publish(topic, args)
}
