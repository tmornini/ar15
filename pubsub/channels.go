package pubsub

import (
	"fmt"
	"time"

	"github.com/tmornini/ar15/wrappers"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// StartChannels create and start a new pubsub using channels
func StartChannels[T event](
	logger wrappers.Logger,
	name string,
	_ T,
) *PubSub[T] {
	ps := &PubSub[T]{
		logger:        logger,
		name:          name,
		publisher:     make(chan T),
		subscriptions: make(chan chan T),
		subscribers:   []chan T{},
		halted:        make(chan halted),
		startedAt:     time.Now(),
	}

	go ps.async()

	return ps
}

// Subscribe to this this publisher
func (ps *PubSub[T]) Subscribe() chan T {
	subscriber := make(chan T)
	ps.subscriptions <- subscriber

	return subscriber
}

// Publish a new event to all subscribers
func (ps *PubSub[T]) Publish(event T) {
	ps.publisher <- event
}

// Stop the new publisher and each of its subscribers
func (ps *PubSub[T]) Close() {
	close(ps.publisher)
}

// Wait for completion of Close() and log statistics
func (ps *PubSub[T]) Wait() {
	<-ps.halted

	ps.logStats()
}

func (ps *PubSub[T]) async() { // in its own goroutine
	var event T
	var ok bool
	var subscriber chan T

LOOP:
	for {
		select {
		case subscriber, ok = <-ps.subscriptions:
			if ok {
				ps.subscribe(subscriber)
			} else {
				ps.logger.Warning("subscription channel closed unexpectedly")
				break LOOP // goroutine exit
			}
		case event, ok = <-ps.publisher:
			if ok {
				ps.publish(event)
				ps.updateStats()
			} else {
				ps.close()
				break LOOP // goroutine exit
			}
		}
	}
}

func (ps *PubSub[T]) subscribe(subscriber chan T) {
	ps.subscribers = append(ps.subscribers, subscriber)
}

func (ps *PubSub[T]) publish(event T) {
	for _, subscriber := range ps.subscribers {
		subscriber <- event
	}
}

func (ps *PubSub[T]) updateStats() {
	ps.eventsIn++
	ps.eventsOut += len(ps.subscribers)
}

func (ps *PubSub[T]) close() {
	close(ps.subscriptions)

	for _, subscriber := range ps.subscribers {
		close(subscriber)
	}

	ps.halted <- halted{}
}

func (ps *PubSub[T]) logStats() {
	elapsed := time.Since(ps.startedAt)

	eventsPerSecond := int(float64(ps.eventsOut) / elapsed.Seconds())

	p := message.NewPrinter(language.English)
	eventsPerSecondWithCommas := p.Sprintf("%d", eventsPerSecond)

	averageLatency := float64(elapsed) / float64(ps.eventsOut)

	ps.logger.Info(
		fmt.Sprintf(
			"pubsub %s stats: events in(%d), out(%d), per second(%s), average latency(%.2fns)",
			ps.name,
			ps.eventsIn,
			ps.eventsOut,
			eventsPerSecondWithCommas,
			averageLatency,
		),
	)
}
