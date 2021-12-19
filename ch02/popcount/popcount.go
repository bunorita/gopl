package popcount

// 8-bitの数(0-255)のそれぞれのpopulation count(the number of 1 bits)
var pc [256]byte

func init() {
	for i := range pc {
		// pc[0] = 0
		// pc[1] = 1
		// pc[2] = 1
		// pc[3] = 2
		// pc[4] = 1
		// ...
		// pc[255] = 8
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	// 64-bit = 8 * 8-bit
	return int(pc[byte(x>>(0*8))] + // for 1-8 bit
		pc[byte(x>>(1*8))] + // for 9-16 bit
		pc[byte(x>>(2*8))] + // ...
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))]) // for 57-64 bit
}

// calculate population count by using loop
func PopCountWithLoop(x uint64) int {
	var c int
	for ; x != 0; x = x >> 1 {
		if x&1 == 1 {
			c++
		}
	}
	return c
}
