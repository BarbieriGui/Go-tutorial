package main

import (
	"fmt"
)

// Operators perform operations on variables and values
/*
+ add
- substracts
/ divide
% returns the division remainder (restante)
++ increment by 1
-- decrement by 1
*/

var n string = ("\n")

func main() {
	var sum1 = 5 + 2 // "+" add two values together
	var sum2 = sum1 + 3
	var sum3 = sum1 + sum2
	fmt.Println(sum3, n)

	var x = 4 // use the assignment operator (=) to assign value
	x += 2    // addition assignment operator (+=) adds a value
	fmt.Println(x, n)

	//List of assignment operators
	// https://www.tutorialspoint.com/go/go_assignment_operators.htm

	// Comparison Operators
	// compare two values and the return is either true (1) or false (0).

	fmt.Println(sum3 == x)
	fmt.Println(sum3 != x)
	fmt.Println(sum3 > x)
	fmt.Println(sum3 < x)
	fmt.Println(sum3 >= x)
	fmt.Println(sum3 <= x)
	fmt.Println()

	// Logical Operators
	// determine the logic between variables or values

	fmt.Println(("x ="), x)
	fmt.Println(x > 5 && x < 10)    // Returns true if both statements are true
	fmt.Println(x == 5 || x < 10)   // Returns true if one of the statements is true
	fmt.Println(!(x > 5 && x < 10)) // Returns false if the result is true
	fmt.Println()

	// Bitwise Operators
	// are used on (binary) numbers

	var a = 9
	var b = 8

	fmt.Printf("a = %b\n", a)
	fmt.Printf("b = %b\n", b)
	fmt.Printf("a&b = %b\n", a&b)     // (AND) Sets each bit to 1 if both bits are 1
	fmt.Printf("a|b = %b\n", a|b)     // (OR) Sets each bit to 1 if one of two bits is 1
	fmt.Printf("a^b = %b\n", a^b)     // (XOR) Sets each bit to 1 if only one of two bits is 1
	fmt.Printf("a<<2 = %b\n", a<<2)   // (0 fill left shift) Shift left by pushing zeros in from the right
	fmt.Printf("a>>2 = %b\n\n", a>>2) // (Signed right shift) Shift right by pushing copies of the leftmost bit in from the left, and let the rightmost bits fall off

}
