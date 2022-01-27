// Package assemblies models assemblies
package assemblies

import "github.com/tmornini/ar15/ammunition"

// Magazine is the interface for magazines
type Magazine interface{}

// NewMagazine returns a new magazine
func NewMagazine() Magazine {
	return &magazine{
		maxCartridges: 30,
	}
}

type magazine struct {
	cartridges    []ammunition.Cartridge
	maxCartridges int
}

// InsertCartridge inserts a cartridge into the magazine
func (m *magazine) InsertCartridge(r ammunition.Cartridge) bool {
	if len(m.cartridges) == m.maxCartridges {
		return false
	}

	m.cartridges = append(m.cartridges, r)

	return true
}

// ExtractCartridge extracts a cartridge from the magazine
func (m *magazine) ExtractCartridge() (bool, ammunition.Cartridge) {
	if len(m.cartridges) == 0 {
		return false, nil
	}

	cartridge, oneFewerCartridge := m.cartridges[len(m.cartridges)-1], m.cartridges[:len(m.cartridges)-1]

	m.cartridges = oneFewerCartridge

	return true, cartridge
}
