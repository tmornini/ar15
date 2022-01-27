package main

import (
	"os"
	"strconv"

	"github.com/tmornini/ar15/pubsub"
	"github.com/tmornini/ar15/wrappers"
)

// EVENTS number of events per publisher
const EVENTS = 1_000

// SUBSCRIBERS number of subscribers per event
const SUBSCRIBERS = 16

type sampleOne struct {
	i int
}

func (e sampleOne) Event() string {
	return "sampleOne " + strconv.Itoa(e.i)
}

type sampleTwo struct {
	s string
}

func (e sampleTwo) Event() string {
	return "sampleTwo " + e.s
}

func main() {
	logger := wrappers.NewNormalAbnormalLogger(os.Stdout, os.Stderr)
	defer logger.Sync()

	// create PubSub instances for integers and strings. All types, including yours, may be used
	sampleOnePubSub := pubsub.Operate(
		logger,
		"sampleOne",
	)

	sampleTwoPubSub := pubsub.Operate(
		logger,
		"sampleTwo",
	)

	// create subscribers in their own goroutines

	for i := 0; i < SUBSCRIBERS; i++ {
		go func() {
			for sampleOne := range sampleOnePubSub.Subscribe() {
				logger.Info(sampleOne.Event())
			}
			logger.Info("sampleOne subscriber stopped")
		}()

		go func() {
			for sampleTwo := range sampleTwoPubSub.Subscribe() {
				logger.Info(sampleTwo.Event())
			}
			logger.Info("sampleTwo subscriber stopped")
		}()
	}

	// create publishers in their own goroutines

	go func() {
		for i := 0; i < EVENTS; i++ {
			sampleOnePubSub.Publish(sampleOne{i})
		}
		sampleOnePubSub.Close()
	}()

	go func() {
		for i := 0; i < EVENTS; i++ {
			sampleTwoPubSub.Publish(sampleTwo{strconv.Itoa(i)})
		}
		sampleTwoPubSub.Close()
	}()

	sampleOnePubSub.Wait()
	sampleTwoPubSub.Wait()
}
