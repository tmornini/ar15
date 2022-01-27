package assemblies

import (
	"github.com/tmornini/ar15/components"
	"github.com/tmornini/ar15/controls"
)

type Lower interface {
	controls.MagazineRelease
	controls.MagazineWell
	controls.Trigger
	controls.Safety

	components.Hammer
}

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
