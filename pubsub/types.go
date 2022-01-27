package pubsub

import (
	"time"

	"github.com/tmornini/ar15/wrappers"
)

type event interface {
	comparable // would prefer to be constrained to channels
}

type halted struct{}

// PubSub is a publisher/subscriber struct
type PubSub[T event] struct {
	logger        wrappers.Logger
	name          string
	lube          int
	publisher     chan T
	subscriptions chan chan T
	subscribers   []chan T
	halted        chan halted
	startedAt     time.Time
	eventsIn      int
	eventsOut     int
}
