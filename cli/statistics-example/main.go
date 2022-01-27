// Package main demonstrates the use of the Statistics type
package main

import (
	"fmt"
	"os"

	"github.com/tmornini/ar15/statistics"
	"github.com/tmornini/ar15/wrappers"
)

func main() {
	logger := wrappers.NewNormalAbnormalLogger(os.Stdout, os.Stderr)
	defer logger.Sync()

	data := []float64{2, 4, 6, 8, 10, 12, 14}

	stats := statistics.NewStatistics("example", logger)

	for _, value := range data {
		stats.AddSample(value)
	}

	fmt.Println(stats)
}
