package ammunition

type Primer interface{}

func NewPrimer() Primer {
	return primer{}
}

type primer struct{}
