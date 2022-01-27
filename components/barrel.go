package components

type Barrel interface{}

func NewBarrel() Barrel {
	return barrel{}
}

type barrel struct {
	length         float32
	inchesPerTwist int
}

func (b barrel) Length() float32 {
	return b.length
}

func (b barrel) InchesPerTwist() int {
	return b.inchesPerTwist
}
