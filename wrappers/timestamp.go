package wrappers

import (
	"time"
)

type Timestamp interface {
	Format(string) string
	IsZero() bool
	RawTime() time.Time
	String() string
	Subtract(other Timestamp) Duration
}

func NewTimestamp() Timestamp {
	return timestamp{time.Now().UTC()}
}

type timestamp struct {
	time.Time
}

func (t timestamp) String() string {
	formatted := t.Time.Format(time.RFC3339Nano)

	if len(formatted) < 27 {
		withoutZ := formatted[0 : len(formatted)-1]

		for i := 0; i < 27-len(withoutZ); i++ {
			withoutZ += "0"
		}

		formatted = withoutZ + "Z"
	}

	return formatted
}

func (t timestamp) Subtract(other Timestamp) Duration {
	return NewDuration(t.Sub(other.RawTime()))
}

func (t timestamp) RawTime() time.Time {
	return t.Time
}
