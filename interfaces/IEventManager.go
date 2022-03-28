package interfaces

import events "github.com/goglue/eventmanager"

type Subscriber events.Subscriber

type IEventManager interface {

	// Attaches a subscriber to an event
	Attach(eventName string, sub events.Subscriber)

	// Dispatches the event across all the subscribers
	Dispatch(eventName string, eventState interface{})

	// Dispatches the event across all the subscribers
	GoDispatch(eventName string, eventState interface{})

	// De attaches a subscriber from an event
	DeAttach(eventName string, subscriber events.Subscriber)
}
