package main

import (
	"fmt"
	"testing/documentation/dog"
)

type myDog struct {
	Name     string
	HumanAge int
}

func main() {
	d1 := myDog{
		Name:     "Rufus",
		HumanAge: 3,
	}
	fmt.Printf("El perrito llamado %s, tiene una edad canina de %d.", d1.Name, dog.Age(d1.HumanAge))
}
