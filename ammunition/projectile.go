package ammunition

type Projectile interface{}

func NewProjectile() Projectile {
	return projectile{}
}

type projectile struct{}
