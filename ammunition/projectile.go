// Package ammunition models ammunition
package ammunition

// Projectile is the interface for projectiles
type Projectile interface{}

// NewProjectile returns a new projectile
func NewProjectile() Projectile {
	return projectile{}
}

type projectile struct{}
