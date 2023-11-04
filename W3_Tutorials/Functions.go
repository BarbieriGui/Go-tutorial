package main

import (
	"fmt"
)

// Naming Rules
// start with letter
// accepts only » (A-z, 0-9, and _ )
// are case-sensitive
// e.g.: myCustomFunction_v2

func myMessage() {
	fmt.Println("Function text")
}

func familyName(fname string, age int) { // "fname" = parameter
	fmt.Println("That's", fname, "Smith,", age, "years old.")
}

func myReturnFunc(x int, y int) int {
	return x + y
}

// can also be done with named return values

func myReturnFunc_2(a int, b int) (result int) {
	result = a + b
	return result // if result not specified, also works
}

func main() {
	myMessage() // call and execute the function
	myMessage()

	// Parameters and Arguments
	// parameters act as variables inside the function (can add as many as needed, just separate them with ",")
	// syntax: func FunctionName(param1 type, param2 type, param3 type) { //code to run } check lines: 17 & 29

	familyName("John", 12) // "John" = argument
	familyName("Marcos", 35)

	fmt.Print("1º Function return: ")
	fmt.Println(myReturnFunc(2, 4))
	fmt.Print("2º Function return: ")
	fmt.Println(myReturnFunc_2(1, 2))

}
