package main

import "fmt"

func main() {
	fmt.Println("1 plus 3 is:", mySum(1, 3))
	fmt.Println("5 plus 4 is:", mySum(5, 4))
	fmt.Println("7 plus 7 is:", mySum(7, 7))
}

func mySum(xi ...int) int {
	var sum int
	for _, v := range xi {
		sum += v
	}
	return sum
}
