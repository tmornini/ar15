// Package weapons models AR15s
package weapons

import (
	"sync"

	"github.com/tmornini/ar15/assemblies"
	"github.com/tmornini/ar15/components"
	"github.com/tmornini/ar15/wrappers"
)

// AR15 is the interface for AR-15s
type AR15 interface {
	assemblies.Upper
	assemblies.Lower

	components.Kinesthetic

	wrappers.Logger

	Operate() error
}

type terminalAR15 struct {
	assemblies.Upper
	assemblies.Lower

	components.Kinesthetic

	wrappers.Logger

	*sync.WaitGroup
}

// NewTerminalAR15 returns a new terminal AR-15
func NewTerminalAR15(l wrappers.Logger, wg *sync.WaitGroup) AR15 {
	return terminalAR15{
		assemblies.NewUpper(),
		assemblies.NewLower(),

		components.NewKinesthetic(),

		l,
		wg,
	}
}

// Operate operates the terminal AR-15
func (t terminalAR15) Operate() error {
	select {}
}
