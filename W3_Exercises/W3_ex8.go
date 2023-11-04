package main

import (
	"fmt"
)

func main() {
	const x int = 50
	const y int = 10
	if x > y {
		fmt.Print("Hello")
	} else if x == y {
		// for equal use ==
		fmt.Print("Bye")
	} else {
		fmt.Print("Ok")
	}
}
