package main

import (
	"fmt"
	"testing/testingBenchmarking/06exercises/secondExercise/quote"
	"testing/testingBenchmarking/06exercises/secondExercise/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
