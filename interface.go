package ar15

import (
	"github.com/tmornini/ar15/assemblies"
	"github.com/tmornini/ar15/components"
	"github.com/tmornini/ar15/wrappers"
)

type Model interface {
	assemblies.Upper
	assemblies.Lower

	components.Kinesthetic

	wrappers.Logger

	Operate() error
}
