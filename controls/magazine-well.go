package controls

type MagazineWell interface{}

func NewMagazineWell() MagazineWell {
	return magazineWell{}
}

type magazineWell struct{}
