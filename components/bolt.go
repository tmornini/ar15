package components

type Bolt interface{}

func NewBolt(position, lockedLimit, batteryLimit float32) Bolt {
	return bolt{}
}

type bolt struct{}
