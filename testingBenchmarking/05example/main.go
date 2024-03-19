package main

import (
	"fmt"
	"strings"
	"testing/testingBenchmarking/05example/mystr"
)

const s = "La vida es un cine"

func main() {
	xs := strings.Split(s, " ")
	for _, v := range xs {
		fmt.Println(v)
	}
	fmt.Println(mystr.Cat(xs))
	fmt.Println(mystr.Join(xs))
}
