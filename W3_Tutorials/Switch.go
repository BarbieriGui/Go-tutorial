package main

import (
	"fmt"
)

// Switch Statement
// Use the switch statement to select one of many code blocks to be executed.
// Only runs the matching statement, does NOT need a BREAK statement.

func main() {

	// Sinle-case switch

	var grade string = ("A")
	switch grade {
	case ("A"):
		fmt.Println("Great job")
	case ("B"):
		fmt.Println("Good job")
	case ("C"):
		fmt.Println("Work harder")
	default:
		fmt.Println("Grade not available") // default is optional | is used if the is no case match
	}

	// Multi-case switch

	var x int = 2
	switch x {
	case 2, 4:
		fmt.Println("odd value")
	case 1, 3:
		fmt.Println("even value")
	default:
		fmt.Println("unkown value")
	}
}
