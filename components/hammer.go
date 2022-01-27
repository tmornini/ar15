package components

type Hammer interface{}

func NewHammer() Hammer {
	return hammer{}
}

type hammer struct{}
