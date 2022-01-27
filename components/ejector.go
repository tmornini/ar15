// Package components models components
package components

// Ejector is the interface for ejector components
type Ejector interface{}

// NewEjector returns a new ejector
func NewEjector() Ejector {
	return ejector{}
}

type ejector struct{}
