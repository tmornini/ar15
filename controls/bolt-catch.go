package controls

type BoltCatch interface{}

func NewBoltCatch() BoltCatch {
	return boltCatch{}
}

type boltCatch struct{}
