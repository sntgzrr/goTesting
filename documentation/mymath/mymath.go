// Package mymath provides Math solutions.
package mymath

// Sum sums an unlimited quantity of int numbers.
func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
