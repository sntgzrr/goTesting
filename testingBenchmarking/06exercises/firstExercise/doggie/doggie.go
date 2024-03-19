// Package doggie leads us know doggie's life.
package doggie

// Age returns human's age into dogs age.
func Age(n int) int {
	return n * 7
}

// AgeTwo returns human's age into dogs age.
func AgeTwo(n int) int {
	count := 0
	for i := 0; i < n; i++ {
		count += 7
	}
	return count
}
