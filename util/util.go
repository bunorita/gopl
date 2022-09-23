package util

import "math"

func IntMin(nums ...int) int {
	min := math.MaxInt // initialize
	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	return min
}
