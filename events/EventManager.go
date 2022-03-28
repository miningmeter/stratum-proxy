package events

import (
	"github.com/goglue/eventmanager"
	"gitlab.com/TitanInd/hashrouter/interfaces"
)

type EventManager struct {
	// interfaces.IEventManager
	eventmanager.Recorder
	*eventmanager.EventManager
}

func NewEventManager() interfaces.IEventManager {
	memory := eventmanager.NewMemoryStorage()
	dispatcher := eventmanager.NewDispatcher()

	newEventManager := &EventManager{}

	newInternalEventManager := eventmanager.NewEventManager(memory, dispatcher, newEventManager)

	newEventManager.EventManager = newInternalEventManager

	return newEventManager
}
