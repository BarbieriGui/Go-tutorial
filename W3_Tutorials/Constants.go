package main

import (
	"fmt"
)

// const CONSTNAME type = value
// const value CANNOT be changed

const PI float32 = 3.14 // Typed constant
const x = 12            // Untyped constants

// Multiple Constants Declaration
const (
	y int    = 12
	i string = "Text"
)

func main() {

	fmt.Println(PI, x)
	fmt.Println(y, i)
}
