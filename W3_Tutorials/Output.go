package main

import (
	"fmt"
)

const a, b string = "Hello", "World"
const c, d string = "Manual", "Space"
const e, f string = "Auto", "Space & Line"

// "\n" creates new lines.
// " " to create space between string arguments.
// %v is used to print the value of the arguments.
// %T is used to print the type of the arguments.

func main() {
	fmt.Print(a, "\n", b) // prints in deafult format.
	fmt.Print("\n", c, " ", d, "\n")
	fmt.Println(e, f) // adds space between arguments and a new line in the end, by default.

	var g string = "Hello"
	var h int = 15

	fmt.Printf("g has value: %v and type: %T\n", g, g)
	fmt.Printf("h has value: %v and type: %T\n", h, h)
}
