// Package controls models controls
package controls

// MagazineWell is the interface for magazine well controls
type MagazineWell interface{}

// NewMagazineWell returns a new magazine well control
func NewMagazineWell() MagazineWell {
	return magazineWell{}
}

type magazineWell struct{}
