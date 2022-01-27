// Package controls models controls
package controls

// ChargingHandle is the interface for charging handle controls
type ChargingHandle interface{}

// NewChargingHandle returns a new charging handle control
func NewChargingHandle() ChargingHandle {
	return chargingHandle{}
}

type chargingHandle struct{}
