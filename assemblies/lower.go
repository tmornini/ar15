// Package assemblies models assemblies
package assemblies

import (
	"github.com/tmornini/ar15/components"
	"github.com/tmornini/ar15/controls"
)

// Lower is the interface for lower assemblies
type Lower interface {
	controls.MagazineRelease
	controls.MagazineWell
	controls.Trigger
	controls.Safety

	components.Hammer
}

// NewLower returns a new lower assembly
func NewLower() Lower {
	return lower{}
}

type lower struct {
	controls.MagazineRelease
	controls.MagazineWell
	controls.Trigger
	controls.Safety

	components.Hammer
}
