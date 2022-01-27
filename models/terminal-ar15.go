package models

import (
	"sync"

	"github.com/tmornini/ar15"
	"github.com/tmornini/ar15/assemblies"
	"github.com/tmornini/ar15/components"
	"github.com/tmornini/ar15/wrappers"
)

type terminalAR15 struct {
	assemblies.Upper
	assemblies.Lower

	components.Kinesthetic

	wrappers.Logger

	*sync.WaitGroup
}

func NewTerminalAR15(l wrappers.Logger, wg *sync.WaitGroup) ar15.Model {
	return terminalAR15{
		assemblies.NewUpper(),
		assemblies.NewLower(),

		components.NewKinesthetic(),

		l,
		wg,
	}
}

func (t terminalAR15) Operate() error {
	select {}
}
