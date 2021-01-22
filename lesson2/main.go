package main

import (
	"fmt"
	"github.com/evgeniyv6/c_algorithms/lesson2/l2"
)

func main() {
	fmt.Printf("Total number of programs from 3 to 20 (recursive algorithm): %d\n", l2.CalcImplementerRecur(3,20, "", ""))
	fmt.Printf("Total number of programs from 3 to 20 (array algorithm): %d\n", l2.CalcImplementerAr(3,20))

	fmt.Printf("Decimal 89 to binary (non recursion): %s\n", l2.DecimalToBin(89))
	fmt.Printf("Decimal 89 to binary (recursion): %d\n", l2.DecimalToBinRecur(89))

	fmt.Printf("Multiply -3 * -10 (recursion): %d\n",l2.MultiplyRecur(-3,-10))
	fmt.Printf("Multiply 3 * -4: %d\n",l2.MultiplyBitOffset(3,-4))


}

