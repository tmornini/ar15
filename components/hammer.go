// Package components models components
package components

// Hammer is the interface for hammer components
type Hammer interface{}

// NewHammer returns a new Hammer component
func NewHammer() Hammer {
	return hammer{}
}

type hammer struct{}
