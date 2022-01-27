// Package controls models controls
package controls

// Trigger is the interface for trigger controls
type Trigger interface{}

// NewTrigger returns a new trigger control
func NewTrigger(position float32, triggerLimit, unpulledLimit float32) Trigger {
	return trigger{
		position:      position,
		triggerLimit:  triggerLimit,
		unpulledLimit: unpulledLimit,
	}
}

type trigger struct {
	position      float32
	unpulledLimit float32
	triggerLimit  float32
}
