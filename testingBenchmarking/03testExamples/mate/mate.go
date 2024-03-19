// Package mate does sum operations.
package mate

// Sum calculate the sum of x amount of numbers.
func Sum(xi ...int) int {
	var s int
	for _, v := range xi {
		s += v
	}
	return s
}
