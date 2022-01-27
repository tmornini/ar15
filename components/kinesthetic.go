package components

type Kinesthetic interface{}

func NewKinesthetic() Kinesthetic {
	return kinesthetic{}
}

type kinesthetic struct{}
