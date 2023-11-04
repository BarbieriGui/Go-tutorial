package main

import (
	"fmt"
)

func main() {
	var day int = 4
	switch day {
	case (1):
		fmt.Print("Saturday")
	case (2):
		fmt.Print("Sunday")
	default:
		fmt.Print("Weekday")
	}
}
