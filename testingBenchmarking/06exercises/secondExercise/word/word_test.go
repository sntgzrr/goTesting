package word

import (
	"fmt"
	"reflect"
	"testing"
)

type testUserCount struct {
	input    string
	expected map[string]int
}

type testCount struct {
	input         string
	expectedCount int
}

func TestUseCount(t *testing.T) {
	t1 := []testUserCount{
		{
			input:    "Hey dude",
			expected: map[string]int{"Hey": 1, "dude": 1},
		},
	}
	for _, v := range t1 {
		r := UseCount(v.input)
		if !reflect.DeepEqual(r, v.expected) {
			t.Error("Expected", v.expected, "Got", r)
		}
	}
}

func TestCount(t *testing.T) {
	t1 := []testCount{
		{
			input:         "Hey dude",
			expectedCount: 2,
		},
	}
	for _, v := range t1 {
		r := Count(v.input)
		if r != v.expectedCount {
			t.Error("Expected", v.expectedCount, "Got", r)
		}
	}
}

func ExampleUseCount() {
	fmt.Println(UseCount("Hey dude"))
	// Output:
	// map[Hey:1 dude:1]
}

func ExampleCount() {
	fmt.Println(Count("Hey dude"))
	// Output:
	// 2
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount("Hey dude")
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count("Hey dude")
	}
}
