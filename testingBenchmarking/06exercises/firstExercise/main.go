package main

import (
	"fmt"
	"testing/testingBenchmarking/06exercises/firstExercise/doggie"
)

type canine struct {
	Name string
	Age  int
}

func main() {
	d1 := canine{
		Name: "Rufus",
		Age:  doggie.Age(10),
	}
	fmt.Println(d1)
	fmt.Println(doggie.AgeTwo(20))
}
