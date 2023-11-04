package main

import (
	"fmt"
)

func main() {
	var x int = 12
	// var := 12 » would let GO decide the type of the variable, and it is not possible to declare it.
	var y float64 = 34.56
	var z string = "words"
	var k bool = false

	fmt.Println(x, y, z, k)

	var student1 string = "John" //type is string
	var student2 = "Jane"        //type is inferred
	h := 2                       //type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(h)
}
