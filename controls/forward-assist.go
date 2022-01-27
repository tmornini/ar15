package controls

type ForwardAssist interface{}

func NewForwardAssist() ForwardAssist {
	return forwardAssist{}
}

type forwardAssist struct{}
