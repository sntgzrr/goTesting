package maths

import (
	"fmt"
	"testing"
)

type test struct {
	input    []int
	expected float64
}

func TestCenteredAverage(t *testing.T) {
	tests := []test{
		{
			input:    []int{1, 2, 3, 4},
			expected: 2.5,
		},
		{
			input:    []int{2, 3, 4, 5, 6},
			expected: 4,
		},
	}
	for _, v := range tests {
		r := CenteredAverage(v.input)
		if r != v.expected {
			t.Error("Expected", v.expected, "Got", r)
		}
	}
}

func ExampleCenteredAverage() {
	fmt.Println(CenteredAverage([]int{3, 4, 5, 6, 7, 8}))
	// Output:
	// 5.5
}

func BenchmarkCenteredAverage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAverage([]int{4, 5, 6, 7, 8, 9, 10})
	}
}
