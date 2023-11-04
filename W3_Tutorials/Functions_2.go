package main

import (
	"fmt"
)

// Store the Return Value in a Variable

func myFunction(x int, y int) (result1 int) {
	result1 = x + y
	return result1
}

// Multiple Return Values

func myFunction2(a int, b string) (result2 int, text1 string) {
	result2 = a + a
	text1 = b + " in stock"
	return
}

// All functions are executed in the main

func main() {
	total := myFunction(3, 2) // variable used to store the return value
	fmt.Println(total)

	fmt.Println(myFunction2(3, "units")) // multiple return values e.g.

	c, d := myFunction2(1, "parts") // storing return values in variables
	fmt.Println(c, d)

	_, e := myFunction2(5, "apples") // omiting the first value
	fmt.Println(e)

	f, _ := myFunction2(5, "apples") // omiting the second value
	fmt.Println(f)
}
