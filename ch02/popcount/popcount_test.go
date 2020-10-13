package popcount_test

import (
	"math"
	"testing"

	"github.com/bunorita/gopl/ch02/popcount"
)

func TestZero(t *testing.T) {
	testZero(t, popcount.PopCount)
	testZero(t, popcount.PopCountWithLoop)
}

func testZero(t *testing.T, popCount func(uint64) int) {
	output := popCount(0)
	if output != 0 {
		t.Errorf("PopCount is %d, but it should be 0", output)
	}
}

func TestAllBits(t *testing.T) {
	testAllBits(t, popcount.PopCount)
	testAllBits(t, popcount.PopCountWithLoop)
}

// test for math.MaxUint64
func testAllBits(t *testing.T, popCount func(uint64) int) {
	output := popCount(math.MaxUint64)
	if output != 64 {
		t.Errorf("PopCount is %d, but it should be 64", output)
	}
}

func TestEachByte(t *testing.T) {
	testEachByte(t, popcount.PopCount)
	testEachByte(t, popcount.PopCountWithLoop)
}

// test for 0xff, 0xff00, 0xff0000, ...
func testEachByte(t *testing.T, popCount func(uint64) int) {
	for i := 0; i < 8; i++ {
		var x uint64 = 0xff << (i * 8)
		output := popCount(x)
		if output != 8 {
			t.Errorf("PopCount is %d, but it should be 8", output)
		}
	}
}

func Test0x555(t *testing.T) {
	test0x5555(t, popcount.PopCount)
	test0x5555(t, popcount.PopCountWithLoop)
}

// test for 0x5555 (= 0b0101010101010101), ...
func test0x5555(t *testing.T, popCount func(uint64) int) {
	for i := 0; i < 4; i++ {
		var x uint64 = 0x5555 << (i * 16)
		output := popCount(x)
		if output != 8 {
			t.Errorf("PopCount is %d, but it should be 8", output)
		}
	}
}

func TestEachOneBit(t *testing.T) {
	testEachOneBit(t, popcount.PopCount)
	testEachOneBit(t, popcount.PopCountWithLoop)
}

func testEachOneBit(t *testing.T, popCount func(uint64) int) {
	for i := 0; i < 64; i++ {
		var x uint64 = 1 << i
		output := popCount(x)
		if output != 1 {
			t.Errorf("PopCount is %d, but it should be 1", output)
		}
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountWithLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountWithLoop(0x1234567890ABCDEF)
	}
}
