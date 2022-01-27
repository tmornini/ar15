// Package components models components
package components

// Bolt is the interface for bolts
type Bolt interface{}

// NewBolt returns a new bolt
func NewBolt() Bolt {
	return bolt{}
}

type bolt struct{}
