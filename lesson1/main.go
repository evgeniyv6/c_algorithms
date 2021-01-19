package main

import (
	"fmt"
	"github.com/evgeniyv6/c_algorithms/lesson1/c_algos"
)

func main() {
	fmt.Println(c_algos.IsAutomorph(90626))

	// try 14 exercise
	fmt.Println("Automorphic numbers:")
	for i := 1; i < 2890625; i++ {
		if c_algos.IsAutomorph(i) {
			fmt.Printf("%d\t", i)
		}
	}
	fmt.Println()
	fmt.Println()

	// try 13 exercise
	fmt.Println("Random numbers:")
	fmt.Printf("Standart rand func: %d\n", c_algos.StandartRand(1, 100))
	fmt.Printf("Lehmer rand func: %d\n", c_algos.KnuthRand(1, 100, 4, 5))

	fmt.Println()
	// try 12 exercise
	fmt.Printf("Max: %d\n", c_algos.MaxFromInt(1, 2, 3))
	fmt.Printf("Min: %d\n", c_algos.MinFromInt(1, -122, -3))
	fmt.Println()

	// try 6 exercise
	for i := 1; i <= 150; i++ {
		age, yearDes := c_algos.YearIdentifier(i)
		fmt.Printf("age - %d  %s # ", age, yearDes)
	}

}
