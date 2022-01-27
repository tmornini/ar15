package controls

type MagazineRelease interface{}

func NewMagazineRelease() MagazineRelease {
	return magazineRelease{}
}

type magazineRelease struct{}
