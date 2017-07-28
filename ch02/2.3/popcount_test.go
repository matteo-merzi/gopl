package popcount

import (
	"testing"
)

// pc[i] is the population count of i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCountTableLoop returns the population count (number of set bits) of x
func PopCountTableLoop(x uint64) int {
	sum := 0
	for i := 0; i < 8; i++ {
		sum += int(pc[byte(x>>uint(i*8))])
	}
	return sum
}

// PopCountTable returns the population count (number of set bits) of x
func PopCountTable(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func bench(b *testing.B, f func(uint64) int) {
	for i := 0; i < b.N; i++ {
		f(uint64(i))
	}
}

func BenchmarkPopCountTable(b *testing.B) {
	bench(b, PopCountTable)
}

func BenchmarkPopCountTableLoop(b *testing.B) {
	bench(b, PopCountTableLoop)
}
