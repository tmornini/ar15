// Package pubsub models a publish/subscribe system
package pubsub

import (
	"time"

	"github.com/tmornini/ar15/statistics"
	"github.com/tmornini/ar15/wrappers"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Event is the interface for events
type Event interface {
	Event() string
}

// PubSub is the interface for pubsubs
type PubSub interface {
	Subscribe() chan Event
	Publish(Event)
	Close()
	Wait()
}

// PubSub is a publisher/subscriber struct
type pubSub struct {
	printer       *message.Printer
	logger        wrappers.Logger
	name          string
	lube          int
	publisher     chan Event
	subscriptions chan chan Event
	subscribers   []chan Event
	halted        chan halted
	statistics    *statistics.Statistics
}

type halted struct{}

// Operate create and start a new pubsub using channels
func Operate(
	logger wrappers.Logger,
	name string,
) PubSub {
	ps := &pubSub{
		printer:       message.NewPrinter(language.English),
		logger:        logger,
		name:          name,
		publisher:     make(chan Event),
		subscriptions: make(chan chan Event),
		subscribers:   []chan Event{},
		halted:        make(chan halted),
		statistics:    statistics.NewStatistics("pubsub "+name+" stats", logger),
	}

	go ps.async()

	return ps
}

// Subscribe to this this publisher
func (ps *pubSub) Subscribe() chan Event {
	subscriber := make(chan Event)
	ps.subscriptions <- subscriber

	return subscriber
}

// Publish a new event to all subscribers
func (ps *pubSub) Publish(e Event) {
	ps.publisher <- e
}

// Stop the new publisher and each of its subscribers
func (ps *pubSub) Close() {
	close(ps.publisher)
}

// Wait for completion of Close() and log statistics
func (ps *pubSub) Wait() {
	<-ps.halted

	ps.log()
}

func (ps *pubSub) subscribe(subscriber chan Event) {
	ps.subscribers = append(ps.subscribers, subscriber)
}

func (ps *pubSub) publish(event Event) {
	for _, subscriber := range ps.subscribers {
		subscriber <- event
	}
}

func (ps *pubSub) updateStatistics(eventStart time.Time) {
	latency := time.Since(eventStart)

	ps.statistics.AddSample(float64(latency.Nanoseconds()))
}

func (ps *pubSub) close() {
	close(ps.subscriptions)

	for _, subscriber := range ps.subscribers {
		close(subscriber)
	}

	ps.halted <- halted{}
}

func (ps *pubSub) async() { // in its own goroutine
	var subscriber chan Event
	var event Event
	var ok bool
	var exit bool

	for {
		eventStart := time.Now()

		if exit {
			break
		}

		select {
		case subscriber, ok = <-ps.subscriptions:
			if ok {
				ps.subscribe(subscriber)
			} else {
				ps.logger.Warning("subscription channel closed unexpectedly")
				exit = true // goroutine exit
			}
		case event, ok = <-ps.publisher:
			if ok {
				ps.publish(event)
				ps.updateStatistics(eventStart)
			} else {
				ps.close()
				exit = true
			}
		}
	}
}

func (ps *pubSub) statisticalMetrics() string {
	return ps.printer.Sprintf(
		"sample size(%d) min (%s) max(%s) mean(%s) standard deviation(%s)",
		ps.statistics.SampleSize(),
		time.Duration(int64(ps.statistics.Min())),
		time.Duration(int64(ps.statistics.Max())),
		time.Duration(int64(ps.statistics.Mean())),
		time.Duration(int64(ps.statistics.StandardDeviation())),
	)
}

func (ps *pubSub) logString() string {
	return ps.printer.Sprintf(
		"%s %s",
		ps.String(),
		ps.statisticalMetrics(),
	)
}

func (ps *pubSub) String() string {
	return "pubSub(" + ps.name + ")"
}

func (ps *pubSub) log() {
	ps.logger.Info(ps.logString())
}
