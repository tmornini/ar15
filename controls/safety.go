// Package controls models controls
package controls

// Safety is the interface for safety controls
type Safety interface{}

// NewSafety returns a new safety control
func NewSafety() Safety {
	return safety{}
}

type safety struct{}
