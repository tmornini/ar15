package controls

type EjectionPortCover interface{}

func NewEjectionPortCover() EjectionPortCover {
	return ejectionPortCover{}
}

type ejectionPortCover struct{}
