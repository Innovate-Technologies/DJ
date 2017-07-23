package at

import (
	"sync"
	"time"

	"github.com/chuckpreslar/emission"
)

// Events is the event emitter to use (first run only)
var Events *emission.Emitter

// At allows you to be able to dispatch events AT certain times
type At struct {
	actions      []Action
	actionsMutex sync.Mutex
	events       *emission.Emitter
}

// Action is a specific event to be send at a specific time
type Action struct {
	Time  time.Time
	Event interface{}
}

var instance *At

// GetInstance gives you a instance of At
func GetInstance() *At {
	if instance == nil {
		instance = &At{
			events:       Events,
			actionsMutex: sync.Mutex{},
		}
		go instance.Run()
	}
	return instance
}

// Run starts checking for the time and emits the events
func (at *At) Run() {
	now := time.Now()
	for {
		at.actionsMutex.Lock()
		for i, action := range at.actions {
			if now.Equal(action.Time) || now.Before(action.Time) {
				go at.events.Emit(action.Event)
				at.actions = append(at.actions[:i], at.actions[i+1:]...)
			}
		}
		at.actionsMutex.Unlock()
		time.Sleep(time.Second)
	}
}

// Add adds an Action to be ran
func (at *At) Add(action Action) {
	at.actionsMutex.Lock()
	at.actions = append(at.actions, action)
	at.actionsMutex.Unlock()
}

// RemoveAllWithEvent removes all actions with a certain event
func (at *At) RemoveAllWithEvent(event interface{}) {
	at.actionsMutex.Lock()
	for i, action := range at.actions {
		if action.Event == event {
			at.actions = append(at.actions[:i], at.actions[i+1:]...)
		}
	}
	at.actionsMutex.Unlock()
}

// RemoveAll removes all events
func (at *At) RemoveAll() {
	at.actionsMutex.Lock()
	at.actions = []Action{}
	at.actionsMutex.Unlock()
}
