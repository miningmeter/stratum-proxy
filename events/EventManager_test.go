package events

import (
	"log"
	"testing"

	"github.com/goglue/eventmanager"
)

type TestSubscriber struct {
	eventmanager.Subscriber
}

var eventPayload string

func (s *TestSubscriber) Update(payload interface{}) {
	log.Println("updating")
	eventPayload = payload.(string)
}

func TestSnapshot(t *testing.T) {
	eventManager := NewEventManager()

	testSubscriber := &TestSubscriber{}

	eventManager.Attach("test", testSubscriber)

	eventManager.Dispatch("test", "tested")

	if eventPayload != "tested" {
		t.Errorf("expected eventPayload to be 'tested';  recieved: %v", eventPayload)
	}
}
