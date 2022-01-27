// Package components models components
package components

// Kinesthetic is the interface for kinesthetic components
type Kinesthetic interface{}

// NewKinesthetic returns a new kinesthetic component
func NewKinesthetic() Kinesthetic {
	return kinesthetic{}
}

type kinesthetic struct{}
