package ammunition

type Cartridge interface {
	Primer
	Projectile
}

func NewCartridge(primer Primer, projectile Projectile) Cartridge {
	return cartridge{primer, projectile}
}

type cartridge struct {
	projectile Projectile
	primer     Primer
}

func (c cartridge) Strike() (Primer, Projectile) {
	return c.primer, c.projectile
}
