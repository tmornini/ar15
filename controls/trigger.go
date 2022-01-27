package controls

type Trigger interface{}

func NewTrigger(position float32, triggerLimit, unpulledLimit float32) Trigger {
	return trigger{
		position:      position,
		triggerLimit:  triggerLimit,
		unpulledLimit: unpulledLimit,
	}
}

type trigger struct {
	position      float32
	unpulledLimit float32
	triggerLimit  float32
}

func (t trigger) Position() float32 {
	return t.position
}

func (t trigger) IsPulled() bool {
	return t.position >= t.unpulledLimit
}

func (t trigger) IsTriggered() bool {
	return t.position >= t.triggerLimit
}
