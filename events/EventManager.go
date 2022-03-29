package events

import (
	"time"

	"github.com/goglue/eventmanager"
	"gitlab.com/TitanInd/hashrouter/interfaces"
)

type EventManager struct {
	eventmanager.Recorder
	*eventmanager.EventManager
	history map[string]map[time.Time]interface{}
}

//TODO: fix nil pointer reference error
func (e *EventManager) Snapshot(event string, payload interface{}, on time.Time) {
	if e.history == nil {
		e.history = make(map[string]map[time.Time]interface{})
	}

	if e.history[event] == nil {
		e.history[event] = make(map[time.Time]interface{})
	}

	e.history[event][on] = payload
}

func NewEventManager() interfaces.IEventManager {
	memory := eventmanager.NewMemoryStorage()
	dispatcher := eventmanager.NewDispatcher()

	newEventManager := &EventManager{}

	newInternalEventManager := eventmanager.NewEventManager(memory, dispatcher, newEventManager)

	newEventManager.EventManager = newInternalEventManager

	return newEventManager
}
