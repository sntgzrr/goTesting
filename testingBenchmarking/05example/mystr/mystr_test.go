package mystr

import (
	"fmt"
	"strings"
	"testing"
)

func TestCat(t *testing.T) {
	s := "Shake not bra"
	xs := strings.Split(s, " ")
	s = Cat(xs)
	if s != "Shake not bra" {
		t.Error("Got", s, "Want", "Shake not bra")
	}
}

func TestJoin(t *testing.T) {
	s := "Shake not bra"
	xs := strings.Split(s, " ")
	s = Join(xs)
	if s != "Shake not bra" {
		t.Error("Got", s, "Want", "Shake not bra")
	}
}

func ExampleCat() {
	s := "Shake not bra"
	xs := strings.Split(s, " ")
	fmt.Println(Cat(xs))
	//Output:
	//Shake not bra
}

func ExampleJoin() {
	s := "Shake not bra"
	xs := strings.Split(s, " ")
	fmt.Println(Join(xs))
	//Output:
	//Shake not bra
}

var s = "Shake not bra"
var xs []string

func BenchmarkCat(b *testing.B) {
	xs = strings.Split(s, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Cat(xs)
	}
}

func BenchmarkJoin(b *testing.B) {
	xs = strings.Split(s, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(xs)
	}
}
