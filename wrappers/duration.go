// Package wrappers wraps the stdlib
package wrappers

import "time"

// Duration is the interface that wraps time.Duration
type Duration interface {
	Plus(other Duration) Duration
	RawDuration() time.Duration
	String() string
}

// NewDuration returns a new Duration
func NewDuration(d time.Duration) Duration {
	return duration{d}
}

type duration struct {
	time.Duration
}

func (d duration) Plus(other Duration) Duration {
	return NewDuration(d.RawDuration() + other.RawDuration())
}

func (d duration) RawDuration() time.Duration {
	return d.Duration
}
