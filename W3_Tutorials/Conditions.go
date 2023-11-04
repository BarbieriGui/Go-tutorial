package main

import (
	"fmt"
)

// can be true or false

var x int = 1
var y int = 2

func main() {
	// suports comparison operations
	fmt.Println(x < y)
	fmt.Println(x <= y)
	fmt.Println(x > y)
	fmt.Println(x >= y)
	fmt.Println(x == y)
	fmt.Println(x != y)
	fmt.Println()
	// suports logical comparison operations
	fmt.Println((x < y) && (x != y)) //"AND"
	fmt.Println((x < y) || (x >= y)) //"OR"
	fmt.Println(!(x != y))           //"NOT"
	fmt.Println()
	// can be applied with
	// if; else; else if; switch
	if x > y {
		fmt.Println("x > y")
	} else if x == y {
		fmt.Println("x == y")
	} else {
		fmt.Println("x != y")
	}
	// switch is similar to a group of else if
	var z int = 10
	switch z {
	case (2):
		fmt.Println("z != 10")
	case (4):
		fmt.Println("z < 10")
	case (10):
		fmt.Println("z = 10")
	}
	fmt.Println()
	// nested if
	var i1 int = 5
	var i2 int = 10
	if i1 < i2 {
		fmt.Println("i1 < than 10")
		if i1 != i2 {
			fmt.Println("Also, i1 != than i2")
			fmt.Println()
		}
	}
}
