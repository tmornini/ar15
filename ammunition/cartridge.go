// Package ammunition models ammunition
package ammunition

// Cartridge combines of a primer and a projectile
type Cartridge interface {
	Primer
	Projectile
}

// NewCartridge returns a new cartridge
func NewCartridge(primer Primer, projectile Projectile) Cartridge {
	return cartridge{primer, projectile}
}

type cartridge struct {
	projectile Projectile
	primer     Primer
}

// Strike returns the post-strike primer and projectile
func (c cartridge) Strike() (Primer, Projectile) {
	return c.primer, c.projectile
}
