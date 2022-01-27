package wrappers

import "time"

type Duration interface {
	Plus(other Duration) Duration
	RawDuration() time.Duration
	String() string
}

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
