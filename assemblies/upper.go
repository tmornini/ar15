// Package assemblies models assemblies
package assemblies

import (
	"github.com/tmornini/ar15/components"
	"github.com/tmornini/ar15/controls"
)

// Upper is the interface for upper assemblies
type Upper interface {
	controls.BoltCatch
	controls.ChargingHandle
	controls.EjectionPortCover
	controls.ForwardAssist
	controls.Safety
	controls.Trigger

	BoltCarrierGroup

	components.Barrel
}

// NewUpper returns a new upper assembly
func NewUpper() Upper {
	return upper{}
}

type upper struct {
	controls.BoltCatch
	controls.ChargingHandle
	controls.EjectionPortCover
	controls.ForwardAssist
	controls.Safety
	controls.Trigger

	BoltCarrierGroup

	components.Barrel
}
