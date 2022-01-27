package components

type FiringPin interface{}

func NewFiringPin() FiringPin {
	return firingPin{}
}

type firingPin struct{}
