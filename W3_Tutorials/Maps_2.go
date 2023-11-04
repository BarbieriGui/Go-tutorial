package main

import (
	"fmt"
)

// Maps Are References
// if 2 variables refer to the same hash, changing one effects the other

func main() {
	var a = map[string]string{"Name": "John", "Age": "32", "Job": "Driver"}
	b := a

	fmt.Println(a)
	fmt.Println(b)

	b["Age"] = "33"
	fmt.Println("After change:")

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println()
	// Iterating Over Maps
	// Maps are unordered structures

	x := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	var y = []string{}
	y = append(y, "one", "two", "three", "four")

	for k, v := range x { // range can return one or two values. In Map (1st value: key, 2nd value: value)
		fmt.Printf("%v : %v ", k, v)
	}

	fmt.Println()

	for _, element := range y { // loop with the defined order
		fmt.Printf("%v : %v, ", element, x[element])
	}
}
