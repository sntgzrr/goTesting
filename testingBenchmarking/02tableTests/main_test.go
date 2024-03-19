package main

import "testing"

func TestMySum(t *testing.T) {
	type test struct {
		data     []int
		response int
	}

	tests := []test{
		test{[]int{2, 4, 6}, 12},
		test{[]int{1, 2, 3}, 6},
		test{[]int{5, 4, 6}, 15},
		test{[]int{2, 4, -1}, 5},
	}

	for _, x := range tests {
		v := mySum(x.data...)
		if v != x.response {
			t.Error("Expected", x.response, "Got", v)
		}
	}
}
