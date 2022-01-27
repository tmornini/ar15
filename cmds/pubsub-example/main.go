package main

import (
	"fmt"
	"os"
	"time"

	"github.com/tmornini/ar15/pubsub"
	"github.com/tmornini/ar15/wrappers"
)

// EVENTS number of events per publisher
const EVENTS = 10_000

// SUBSCRIBERS number of subscribers per event
const SUBSCRIBERS = 16

func main() {
	logger := wrappers.NewNormalAbnormalLogger(os.Stdout, os.Stderr)
	defer logger.Sync()

	// create PubSub instances for integers and strings. All types, including yours, may be used
	integerPubSub := pubsub.StartChannels(
		logger,
		"integer",
		1, // type of the publisher is derived by the type of this argument which is otherwise unused
	)

	stringPubSub := pubsub.StartChannels(
		logger,
		"string",
		"", // the type of the publisher is derived from this argument which is otherwise unused
	)

	// create subscribers in their own goroutines

	for i := 0; i < SUBSCRIBERS; i++ {
		go func() {
			for range integerPubSub.Subscribe() {
				// fmt.Printf("i%d", event)
			}
			logger.Info("integer subscriber stopped")
		}()

		go func() {
			for range stringPubSub.Subscribe() {
				// fmt.Printf("s%s", event)
			}
			logger.Info("string subscriber stopped")
		}()
	}

	time.Sleep(10 * time.Millisecond)

	// create publishers in their own goroutines

	go func() {
		for i := 0; i < EVENTS; i++ {
			integerPubSub.Publish(i)
		}
		integerPubSub.Close()
	}()

	go func() {
		for i := 0; i < EVENTS; i++ {
			stringPubSub.Publish(fmt.Sprintf("%d", i))
		}
		stringPubSub.Close()
	}()

	integerPubSub.Wait()
	stringPubSub.Wait()
}
