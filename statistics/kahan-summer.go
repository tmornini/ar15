package statistics

// KahanSummer implements Kahan's summation algorithm.
type KahanSummer struct {
	sum          float64
	compensation float64
}

// NewKahanSummer returns a new KahanSummer.
func NewKahanSummer() *KahanSummer {
	return new(KahanSummer)
}

// Add adds a new float64 to the KahanSummer.
func (ks *KahanSummer) Add(value float64) {
	compensatedValue := value - ks.compensation
	nextSum := ks.sum + compensatedValue
	ks.compensation = (nextSum - ks.sum) - compensatedValue
	ks.sum = nextSum
}

// Sum returns the current sum.
func (ks *KahanSummer) Sum() float64 {
	return ks.sum
}
