// Package controls models controls
package controls

// EjectionPortCover is the interface for ejection port cover controls
type EjectionPortCover interface{}

// NewEjectionPortCover returns a new ejection port cover control
func NewEjectionPortCover() EjectionPortCover {
	return ejectionPortCover{}
}

type ejectionPortCover struct{}
