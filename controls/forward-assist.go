// Package controls models controls
package controls

// ForwardAssist is the interface for forward assist controls
type ForwardAssist interface{}

// NewForwardAssist returns a new forward assist control
func NewForwardAssist() ForwardAssist {
	return forwardAssist{}
}

type forwardAssist struct{}
