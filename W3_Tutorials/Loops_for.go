package main

import "fmt"

// For Loops
// loops through a block of code a specified number of times. (only option of loop in GO)

func main() {
	for i := 1; i < 10; i++ {
		if i == 3 {
			continue //  skip one or more iterations in the loop.
		} else if i == 6 {
			break // used to break/terminate the loop execution
		}
		fmt.Println(i)
	}

	fmt.Println()

	// Nested Loops

	color := [2]string{"blue", "green"}
	toy := [3]string{"soldier", "car", "plane"}
	for c := 0; c < len(color); c++ {
		for t := 0; t < len(toy); t++ {
			fmt.Println(color[c], toy[t])
		}
	}

	fmt.Println()

	// Range keyword
	// is used to more easily iterate over an array, slice or map. It returns both the index and the value.
	// syntax: for index, value := array|slice|map { code to be executed for each iteration }

	grades := [4]string{"A", "B", "C", "D"}
	for idx, val := range grades {
		fmt.Printf("%v\t%v\n", idx, val) // "\t" = tab space
	}

	fmt.Println()

	// to omit the idx or val, replace it by (_)

	for _, val := range grades {
		fmt.Printf("%v\t", val)
	}

}
