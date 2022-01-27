// Package controls models controls
package controls

// MagazineRelease is the interface for magazine release controls
type MagazineRelease interface{}

// NewMagazineRelease returns a new magazine release control
func NewMagazineRelease() MagazineRelease {
	return magazineRelease{}
}

type magazineRelease struct{}
