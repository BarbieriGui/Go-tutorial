package main

import (
	"fmt"
)

// Recursion Functions
// A func that calls it self and reaches a stop condition

func testcount(x int) int {
	if x == 6 {
		return 0
	}
	fmt.Println(x)
	return testcount(x + 1)
}

func factorial_recursion(x float64) (y float64) {
	if x > 0 {
		y = x * factorial_recursion(x-1)
	} else {
		y = 1
	}
	return
}

func main() {
	testcount(1)
	fmt.Println()
	fmt.Println(factorial_recursion(4))
}
