// Package mate can generate example outputs in our documentation code.
package mate

import "fmt"

func ExampleSum() {
	fmt.Println(Sum(1, 2, 3, 4))
	// Output:
	// 10
}
