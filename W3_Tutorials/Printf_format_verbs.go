package main

import (
	"fmt"
)

func main() {

	var i = 15.5
	var txt = "Hello Friends!"

	fmt.Printf("%v\n", i)   // Prints the value in the default format
	fmt.Printf("%#v\n", i)  // Prints the value in Go-syntax format
	fmt.Printf("%v%%\n", i) // Prints the % sign
	fmt.Printf("%T\n", i)   // Prints the type of the value

	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n\n", txt)

	// Integer Formatting Verbs

	var x = 60

	fmt.Printf("%b\n", x)     // Base 2
	fmt.Printf("%d\n", x)     // Base 10
	fmt.Printf("%+d\n", x)    // Base 10 and always show sign
	fmt.Printf("%o\n", x)     // Base 8
	fmt.Printf("%O\n", x)     // Base 8, with leading 0o
	fmt.Printf("%x\n", x)     // Base 16, lowercase
	fmt.Printf("%X\n", x)     // Base 16, uppercase
	fmt.Printf("%#x\n", x)    // Base 16, with leading 0x
	fmt.Printf("%4d\n", x)    // Pad with spaces (width 4, right justified)
	fmt.Printf("%-4d\n", x)   // Pad with spaces (width 4, left justified)
	fmt.Printf("%04d\n\n", x) // Pad with zeroes (width 4)

	// String Formatting Verbs

	var txtNew = "Hi"

	fmt.Printf("%s\n", txtNew)    // Prints the value as plain string
	fmt.Printf("%q\n", txtNew)    // Prints the value as a double-quoted string
	fmt.Printf("%8s\n", txtNew)   // Prints the value as plain string (width 8, right justified)
	fmt.Printf("%-8s\n", txtNew)  // Prints the value as plain string (width 8, left justified)
	fmt.Printf("%x\n", txtNew)    // Prints the value as hex dump of byte values
	fmt.Printf("% x\n\n", txtNew) // Prints the value as hex dump with spaces

	// Boolean Formatting Verbs

	var t bool = true
	var f bool = false

	fmt.Printf("%t\n", t)
	fmt.Printf("%t\n\n", f)

	// Float Formatting Verbs

	var y = 3.141

	fmt.Printf("%e\n", y)    // Scientific notation with 'e' as exponent
	fmt.Printf("%f\n", y)    // Decimal point, no exponent
	fmt.Printf("%.2f\n", y)  // Default width, precision 2 » can change precision value
	fmt.Printf("%6.2f\n", y) // Width 6, precision 2 » can change width and precision value
	fmt.Printf("%g\n", y)    // Exponent as needed, only necessary digits
}
