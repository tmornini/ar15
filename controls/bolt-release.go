package controls

type BoltRelease interface{}

func NewBoltRelease() BoltRelease {
	return boltRelease{}
}

type boltRelease struct{}
