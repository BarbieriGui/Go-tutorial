package main

import (
	"fmt"
)

func main() {
	var a bool = true
	var b int = 10
	var c float64 = 12.45
	var d string = "Horray!"

	fmt.Println("Boolean: ", a)
	fmt.Println("Integer: ", b)
	fmt.Println("Float: ", c)
	fmt.Println("String: ", d)

	// Boolean Data Type

	var b1 bool = true
	var b2 = true
	var b3 bool
	b4 := true

	var x string = "\n"
	fmt.Print(b1, x, b2, x, b3, x, b4, x)

	// Integer Data Type
	/* The default type is int
	Signed integers - can store + & - values
	Unsigned integers - can only store non-negativa values
	*/

	var a1 int = 50
	var a2 int = -50
	var a3 uint = 50
	fmt.Printf("Type: %T, value: %v\n", a1, a1)
	fmt.Printf("Type: %T, value: %v\n", a2, a2)
	fmt.Printf("Type: %T, value: %v\n", a3, a3)

	//Float Data Type
	//float32 = 32 bits
	//float64 = 64 bits (default)

	var f1 float32 = 3.4e+38  //max value
	var f2 float64 = 1.7e+308 //max value
	fmt.Printf("Type: %T, value: %v\n", f1, f1)
	fmt.Printf("Type: %T, value: %v\n", f2, f2)

	//String Data Type
	//String data is used to store text, ("text")

	var txt1 string = "Hi"
	var txt2 string
	txt3 := "Hello"

	fmt.Printf("Type: %T, value: %v\n", txt1, txt1)
	fmt.Printf("Type: %T, value: %v\n", txt2, txt2)
	fmt.Printf("Type: %T, value: %v\n", txt3, txt3)
}
