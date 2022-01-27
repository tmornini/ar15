package assemblies

import "github.com/tmornini/ar15/components"

type BoltCarrierGroup interface {
	components.Ejector
	components.Extractor
	components.FiringPin
}

func NewBoltCarrierGroup() BoltCarrierGroup {
	return boltCarrierGroup{}
}

type boltCarrierGroup struct {
	components.Ejector
	components.Extractor
	components.FiringPin
}
