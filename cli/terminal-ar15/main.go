// Package main demonstrates the use of the TerminalAR15 type.
package main

import (
	"os"
	"sync"

	"github.com/tmornini/ar15/weapons"
	"github.com/tmornini/ar15/wrappers"
)

func main() {
	log := wrappers.NewNormalAbnormalLogger(os.Stdout, os.Stderr)
	wg := &sync.WaitGroup{}

	weapon := weapons.NewTerminalAR15(log, wg)

	err := weapon.Operate()

	if err != nil {
		log.Halt(err)
	}

	log.Info("weapon exited successfully")
}
