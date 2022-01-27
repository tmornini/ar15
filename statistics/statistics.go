package statistics

import (
	"fmt"
	"math"

	"github.com/tmornini/ar15/wrappers"
)

// Statistics is a struct that stores the current mean, variance, count min and max.
type Statistics struct {
	name   string
	logger wrappers.Logger

	sampleSize float64
	min        float64
	max        float64

	mean  KahanSummer
	sigma KahanSummer
}

// NewStatistics returns a new Statistics which implements Welford's online algorithm
func NewStatistics(name string, logger wrappers.Logger) *Statistics {
	return &Statistics{
		name:   name,
		logger: logger,
		min:    math.MaxFloat64,
	}
}

// AddSample adds a new value to the current Statistics's algorithm state.
func (stats *Statistics) AddSample(value float64) {
	stats.sampleSize++

	lastMean := stats.Mean()

	stats.mean.Add(
		(value - lastMean) / stats.sampleSize,
	)

	stats.sigma.Add(
		(value - lastMean) * (value - stats.Mean()),
	)

	if value < stats.min {
		stats.min = value
	} else if value > stats.max {
		stats.max = value
	}
}

// SampleSize returns the current count.
func (stats *Statistics) SampleSize() int {
	return int(stats.sampleSize)
}

// Mean returns the current mean.
func (stats *Statistics) Mean() float64 {
	return stats.mean.Sum()
}

// Sigma returns the current sigma.
func (stats *Statistics) Sigma() float64 {
	return stats.sigma.Sum()
}

// Min returns the current min.
func (stats *Statistics) Min() float64 {
	return stats.min
}

// Max returns the current max.
func (stats *Statistics) Max() float64 {
	return stats.max
}

// Variance returns the current variance.
func (stats *Statistics) Variance() float64 {
	if stats.sampleSize == 0 {
		panic("variance require at least one sample")
	}

	return stats.sigma.Sum() / stats.sampleSize
}

// SampleVariance returns the current sample variance.
func (stats *Statistics) SampleVariance() float64 {
	if stats.sampleSize < 2 {
		panic("sample variance require at least two samples")
	}

	return stats.sigma.Sum() / (stats.sampleSize - 1)
}

// StandardDeviation returns the current standard deviation.
func (stats *Statistics) StandardDeviation() float64 {
	if stats.sampleSize == 0 {
		panic("variance require at least one sample")
	}

	return math.Sqrt(stats.Variance())
}

// String returns a string representation of the current Statistics.
func (stats *Statistics) String() string {
	return fmt.Sprintf(
		"%s: sample size(%d) min(%.2f) max(%.2f) mean(%.2f) standard deviation(%.2f)",
		stats.name,
		stats.SampleSize(),
		stats.Min(),
		stats.Max(),
		stats.Mean(),
		stats.StandardDeviation(),
	)
}

// Log logs the current Statistics.
func (stats *Statistics) Log() {
	stats.logger.Info(stats.String())
}
