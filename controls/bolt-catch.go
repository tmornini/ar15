// Package controls models controls
package controls

// BoltCatch is the interface for bolt catch controls
type BoltCatch interface{}

// NewBoltCatch returns a new bolt catch control
func NewBoltCatch() BoltCatch {
	return boltCatch{}
}

type boltCatch struct{}
