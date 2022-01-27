// Package controls models controls
package controls

// BoltRelease is the interface for bolt release controls
type BoltRelease interface{}

// NewBoltRelease returns a new bolt release control
func NewBoltRelease() BoltRelease {
	return boltRelease{}
}

type boltRelease struct{}
