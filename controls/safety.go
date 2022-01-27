package controls

type Safety interface{}

func NewSafety() Safety {
	return safety{}
}

type safety struct{}
