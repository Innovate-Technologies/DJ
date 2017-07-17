package cron

import (
	"strconv"
	"sync"
	"time"

	"github.com/chuckpreslar/emission"
)

// Events is the event emitter to use (first run only)
var Events *emission.Emitter

// Cron allows events to be sent out like cron jobs
type Cron struct {
	actions      []Action
	actionsMutex sync.Mutex
	events       *emission.Emitter
}

// Action is a specific event to be send at the matching times
type Action struct {
	Minute     string // to allow *
	Hour       string
	DayOfMonth string
	Month      string
	DayOfWeek  string // Sunday is 7 here!
	Event      interface{}
}

var instance *Cron
var once sync.Once

// GetInstance gives you a instance of At
func GetInstance() *Cron {
	once.Do(func() {
		instance = &Cron{
			events:       Events,
			actionsMutex: sync.Mutex{},
		}
		go instance.Run()
	})
	return instance
}

// Run starts checking for the time and emits the events
func (cron *Cron) Run() {
	now := time.Now()
	for {
		cron.actionsMutex.Lock()
		for _, action := range cron.actions {
			needsExecution := true
			if i, err := strconv.ParseInt(action.Minute, 10, 64); err == nil && i != int64(now.Minute()) {
				needsExecution = false
			}
			if i, err := strconv.ParseInt(action.Hour, 10, 64); err == nil && i != int64(now.Hour()) {
				needsExecution = false
			}
			if i, err := strconv.ParseInt(action.DayOfMonth, 10, 64); err == nil && i != int64(now.Day()) {
				needsExecution = false
			}
			if i, err := strconv.ParseInt(action.Month, 10, 64); err == nil && i != int64(now.Month()) {
				needsExecution = false
			}

			day := int64(now.Weekday())
			if day == 0 {
				day = 7
			}
			if i, err := strconv.ParseInt(action.DayOfWeek, 10, 64); err == nil && i != day {
				needsExecution = false
			}

			if needsExecution {
				cron.events.Emit(action.Event)
			}
		}
		cron.actionsMutex.Unlock()
		time.Sleep(time.Second)
	}
}

// Add adds an Action to be ran
func (cron *Cron) Add(action Action) {
	cron.actionsMutex.Lock()
	cron.actions = append(cron.actions, action)
	cron.actionsMutex.Unlock()
}

// RemoveAllWithEvent removes all actions with a certain event
func (cron *Cron) RemoveAllWithEvent(event interface{}) {
	cron.actionsMutex.Lock()
	for i, action := range cron.actions {
		if action.Event == event {
			cron.actions = append(cron.actions[:i], cron.actions[i+1:]...)
		}
	}
	cron.actionsMutex.Unlock()
}

// RemoveAll removes all events
func (cron *Cron) RemoveAll() {
	cron.actionsMutex.Lock()
	cron.actions = []Action{}
	cron.actionsMutex.Unlock()
}
