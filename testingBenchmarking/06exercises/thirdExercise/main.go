package main

import (
	"fmt"
	"testing/testingBenchmarking/06exercises/thirdExercise/maths"
)

func main() {
	x := gen()
	for _, v := range x {
		fmt.Println(maths.CenteredAverage(v))
	}
}

func gen() [][]int {
	a := []int{1, 2, 3, 4}
	b := []int{2, 3, 4, 5, 6}
	c := []int{3, 4, 5, 6, 7, 8}
	d := []int{4, 5, 6, 7, 8, 9, 10}
	e := []int{5, 6, 7, 8, 9, 10, 11, 12}
	f := [][]int{a, b, c, d, e}
	return f
}
