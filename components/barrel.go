// Package components models components
package components

// Barrel is the interface for barrels
type Barrel interface{}

// NewBarrel returns a new barrel
func NewBarrel() Barrel {
	return barrel{}
}

type barrel struct {
	length         float32
	inchesPerTwist int
}
