package event

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

// Dispatcher ..
type Dispatcher struct {
	jobs   chan job
	events map[Name]Listener
}

// NewDispatcher ..
func NewDispatcher(buffCapacity int) *Dispatcher {
	d := &Dispatcher{
		jobs:   make(chan job, buffCapacity),
		events: make(map[Name]Listener),
	}

	go d.consume()

	return d
}

// Register ..
func (d *Dispatcher) Register(listener Listener, names []Name) error {
	for _, name := range names {
		if _, ok := d.events[name]; ok {
			return fmt.Errorf("the '%s' event is already registered", name)
		}
		log.Info("REGISTERED EVENT NAME => ", name)
		d.events[name] = listener
	}

	return nil
}

// Dispatch ..
func (d *Dispatcher) Dispatch(name Name, event interface{}) {
	if _, ok := d.events[name]; !ok {
		fmt.Printf("the '%s' event is not registered\n", name)
		return
	}
	log.Println("[EVENT CONSUMED] => ", name)
	d.jobs <- job{eventName: name, eventType: event}
}

// consume ..
func (d *Dispatcher) consume() {
	for job := range d.jobs {
		d.events[job.eventName].Listen(job.eventType)
	}
}
