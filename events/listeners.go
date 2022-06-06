package events

import (
	"fmt"
	"gitlab-webhook/internal/libraries/event"
	"reflect"
)

// Listener ..
type Listener struct {
}

var dispatcher *event.Dispatcher

func init() {

	eventsName := []event.Name{
		WebhookReceived,

		RunCommandSucceeded,
		RunCommandFailed,
	}
	dispatcher = event.NewDispatcher(len(eventsName))

	if err := dispatcher.Register(Listener{}, eventsName); err != nil {
		fmt.Println("Error while registering events")
	}
}

func Dispatch(name event.Name, event interface{})  {
	go dispatcher.Dispatch(name, event)
}

// Listen ..
func (u Listener) Listen(event interface{}) {
	defer handleFromPanic(event)

	switch eventDispatcher := event.(type) {

	case WebhookReceivedEvent:
		go eventDispatcher.RunCommand()

	case RunCommandSucceededEvent:
		eventDispatcher.SendToDiscord()

	case RunCommandFailedEvent:
		eventDispatcher.SendToDiscord()
	}
}

// handleFromPanic ..
func handleFromPanic(event interface{}) {
	if r := recover(); r != nil {
		fmt.Printf("Recovering from panic in : %v \n ==>> %v \n", reflect.TypeOf(event).Name(), r)
	}
}
