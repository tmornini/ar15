// Package wrappers wraps the stdlib
package wrappers

import (
	"fmt"
	"time"
)

// Timer is the interface that wraps stdlib timing
type Timer interface {
	RecordIteration(Timestamp)
	String() string
}

// NewTimer returns a new Timer
func NewTimer(name string) Timer {
	return &timer{
		name:                    name,
		lastIterationRecordedAt: timestamp{time.Time{}},
		totalDuration:           duration{0},
	}
}

type timer struct {
	name                    string
	count                   int64
	lastIterationRecordedAt Timestamp
	thisDuration            Duration
	totalDuration           Duration
}

// RecordIteration records an iteration
func (t *timer) RecordIteration(now Timestamp) {
	var lastIterationRecordedAt Timestamp

	if t.lastIterationRecordedAt.IsZero() {
		lastIterationRecordedAt = now
	} else {
		lastIterationRecordedAt = t.lastIterationRecordedAt
	}

	t.thisDuration = now.Subtract(lastIterationRecordedAt)
	t.totalDuration = t.totalDuration.Plus(t.thisDuration)

	t.lastIterationRecordedAt = now

	if lastIterationRecordedAt == now {
		return
	}

	t.count++
}

// String returns a string representation of the timer
func (t timer) String() string {
	if t.count == 0 {
		return fmt.Sprintf(
			"TIMER: %s created...",
			t.name,
		)
	}

	return fmt.Sprintf(
		"TIMER: %25s, %10s %10s avg",
		t.name,
		t.thisDuration.String(),
		t.averageDuration().String(),
	)
}

// averageDuration returns the average duration
func (t timer) averageDuration() Duration {
	if t.count == 0 {
		return NewDuration(0)
	}

	return NewDuration(
		time.Duration(
			int64(
				float32(t.totalDuration.RawDuration())/float32(t.count),
			) / 1000 * 1000,
		),
	)
}
