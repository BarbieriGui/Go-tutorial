package main

import (
	"fmt"
)

var a string // default value ""
var b int
var c bool
var e, f, g int = 10, 20, 30
var h, i = 4.12, "Test"
var (
	j float64 = 21.31
	k string  = "Group"
	l bool    = true
)

func main() {

	// a = 1 » could define the value latter
	// h, i := 4, "Test" » could do like this also
	d := "d" // can only be used inside the func

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e, f, g)
	fmt.Println(h, i)
	fmt.Println(j)
	fmt.Println(k, l)
}
