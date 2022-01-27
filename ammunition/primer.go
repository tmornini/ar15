// Package ammunition models ammunition
package ammunition

// Primer is the interface for primers
type Primer interface{}

// NewPrimer returns a new primer
func NewPrimer() Primer {
	return primer{}
}

type primer struct{}
