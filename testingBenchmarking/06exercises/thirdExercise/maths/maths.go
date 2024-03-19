// Package maths makes math operations.
package maths

import "sort"

// CenteredAverage calculates the centered average of a slice of int.
func CenteredAverage(xi []int) float64 {
	sort.Ints(xi)
	xi = xi[1:(len(xi) - 1)]
	sum := 0
	for _, v := range xi {
		sum += v
	}
	f := float64(sum) / float64(len(xi))
	return f
}
