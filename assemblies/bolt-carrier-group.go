// Package assemblies models assemblies
package assemblies

import "github.com/tmornini/ar15/components"

// BoltCarrierGroup is the interface for bolt carrier groups
type BoltCarrierGroup interface {
	components.Ejector
	components.Extractor
	components.FiringPin
}

// NewBoltCarrierGroup returns a new bolt carrier group
func NewBoltCarrierGroup() BoltCarrierGroup {
	return boltCarrierGroup{}
}

type boltCarrierGroup struct {
	components.Ejector
	components.Extractor
	components.FiringPin
}
