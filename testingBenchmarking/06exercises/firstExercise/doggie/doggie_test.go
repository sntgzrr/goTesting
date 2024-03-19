package doggie

import (
	"fmt"
	"testing"
)

func TestAge(t *testing.T) {
	age := Age(3)
	if age != 21 {
		t.Error("Expected", 21, "Got", age)
	}
}

func TestAgeTwo(t *testing.T) {
	age := AgeTwo(3)
	if age != 21 {
		t.Error("Expected", 21, "Got", age)
	}
}

func ExampleAge() {
	fmt.Println(Age(4))
	// Output:
	//28
}

func ExampleAgeTwo() {
	fmt.Println(AgeTwo(4))
	// Output:
	//28
}

func BenchmarkAge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Age(3)
	}
}

func BenchmarkAgeTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AgeTwo(3)
	}
}
