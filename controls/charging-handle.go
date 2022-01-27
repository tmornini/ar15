package controls

type ChargingHandle interface{}

func NewChargingHandle() ChargingHandle {
	return chargingHandle{}
}

type chargingHandle struct{}
