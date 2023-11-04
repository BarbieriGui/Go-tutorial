package main

import (
	"fmt"
)

const n = "\n"

func main() {

	var arr1 = [3]string{"Maria", "João", "Marcos"}
	arr2 := [5]int{4, 5, 6, 7, 8}
	var arr3 = [...]int{1, 2, 3}
	arr4 := [...]int{4, 5, 6, 7, 8}

	fmt.Print(arr1, n, arr2, n, arr3, n, arr4, n)
	fmt.Println(arr3[0], arr3[2]) // array index starts at 0
	arr1[0] = "Teresa"            // replaces "Maria"
	arr2[3] = 9                   // replaces 7
	fmt.Println(arr1)
	fmt.Print(arr2, n, n)

	// Array initialization
	// init = dafault value is 0
	// string = default value is ""

	arr5 := [3]int{}             //not initialized
	arr6 := [3]int{1, 2}         //partially initialized
	arr7 := [3]int{1, 2, 3}      //fully initialized
	arr8 := [5]int{1: 10, 2: 40} //assigned value 10 to array index 1 and value 40 to array index 2
	arr9 := [1]string{}          //not initialized
	fmt.Println(arr5)
	fmt.Println(arr6)
	fmt.Println(arr7)
	fmt.Println(arr8)
	fmt.Println(arr9)
	fmt.Println(len(arr7)) //(len(array name)) can be used to find the array's length
	fmt.Println(len(arr8))
}
