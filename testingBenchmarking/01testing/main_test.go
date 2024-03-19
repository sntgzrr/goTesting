package main

import "testing"

func TestMySum(t *testing.T) {
	v := mySum(2, 8)
	if v != 10 {
		t.Error("Expected", 10, "Got", v)
	}
}
