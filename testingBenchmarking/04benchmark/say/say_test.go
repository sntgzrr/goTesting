package say

import (
	"fmt"
	"testing"
)

// TestSay do a test of Say function.
func TestSay(t *testing.T) {
	s := Say("Santiago")
	if s != "Welcome Santiago" {
		t.Error("Expected", "Welcome Santiago", "Got", s)
	}
}

// ExampleSay shows an example in our documentation using Say function.
func ExampleSay() {
	fmt.Println(Say("Santiago"))
	//Output:
	//Welcome Santiago
}

// BenchmarkSay allows us to measure the speed of execution.
func BenchmarkSay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Say("Santiago")
	}
}
