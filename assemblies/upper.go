package assemblies

import (
	"github.com/tmornini/ar15/components"
	"github.com/tmornini/ar15/controls"
)

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
