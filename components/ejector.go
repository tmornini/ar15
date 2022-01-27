package components

type Ejector interface{}

func NewEjector() Ejector {
	return ejector{}
}

type ejector struct{}
