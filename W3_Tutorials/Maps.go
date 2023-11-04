package main

import (
	"fmt"
)

// Go Maps are used to store data values in key:value pairs
// it is unordered, changeable and does not allow duplicates
// dafult value = nil
/* Syntax
var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
b := map[KeyType]ValueType{key1:value1, key2:value2,...}

Can create using make() » Syntax
	var a = make(map[KeyType]ValueType)
	b := make(map[KeyType]ValueType)		*/

func main() {
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2}

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)
	fmt.Println()

	var a2 = make(map[string]string)
	a2["brand"] = "Ford"
	a2["model"] = "Mustang"
	a2["year"] = "1964"

	b2 := make(map[string]int)
	b2["Oslo"] = 1
	b2["Bergen"] = 2

	fmt.Printf("a2\t%v\n", a2)
	fmt.Printf("b2\t%v\n", b2)
	fmt.Println()

	// Accessing Map Elements
	// Syntax » value = map_name[key]

	fmt.Printf(a2["model"])
	fmt.Println()

	// Updating and Adding Map Elements
	// Syntax map_name[key] = value

	fmt.Println(a2)
	a2["year"] = "1968"
	a2["color"] = "black"
	fmt.Println(a2)

	// Remove Element
	// Syntax delete(map_name, key)

	delete(a2, "year")
	fmt.Println(a2)

	// Creating an Empty Map
	// the correct way is using the make() function

	var c = make(map[string]string)
	var d map[string]string

	fmt.Print("\n", c == nil)
	fmt.Printf("\t%v\n", c)
	fmt.Print(d == nil)
	fmt.Printf("\t%v\n", d)

	// Check for specific elements in Maps
	// Syntax val, ok := map_name[key]

	var x = map[string]string{"version": "V2.7", "year": "2007", "owner": ""}

	val1, ok1 := x["version"]
	val2, ok2 := x["owner"]
	val3, ok3 := x["model"]
	_, ok4 := x["year"]

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
	fmt.Println(ok4)
}
