// Package components models components
package components

// FiringPin is the interface for firing pin components
type FiringPin interface{}

// NewFiringPin returns a new firing pin
func NewFiringPin() FiringPin {
	return firingPin{}
}

type firingPin struct{}
