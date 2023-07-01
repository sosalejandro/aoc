package helpers

import "sort"

// MinInt returns the smallest value in a slice of ints
func MinInt(values ...int) int {
	// Sort values and return first index
	sort.Ints(values)
	smallest := values[0]
	return smallest
}
