package events

import (
	"fmt"
	"time"

	"github.com/goglue/eventmanager"
	"gitlab.com/TitanInd/hashrouter/interfaces"
)

type EventManager struct {
	*eventmanager.EventManager
	recorder *Recorder
}

type Recorder struct {
	eventmanager.Recorder
	history map[string]map[time.Time]interface{}
}

//TODO: fix nil pointer reference error
func (r *Recorder) SnapShot(event string, payload interface{}, on time.Time) {
	fmt.Println("snapshot")
	if r.history == nil {
		r.history = make(map[string]map[time.Time]interface{})
	}

	if r.history[event] == nil {
		r.history[event] = make(map[time.Time]interface{})
	}

	r.history[event][on] = payload
}

func NewEventManager() interfaces.IEventManager {
	memory := eventmanager.NewMemoryStorage()
	dispatcher := eventmanager.NewDispatcher()

	newEventManager := &EventManager{
		recorder: &Recorder{},
	}

	newInternalEventManager := eventmanager.NewEventManager(memory, dispatcher, newEventManager.recorder)

	newEventManager.EventManager = newInternalEventManager

	return newEventManager
}
