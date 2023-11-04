package main

import (
	"fmt"
)

var n = ("\n")

// Slices are a more powerfull version of arrays but unlike arrays, the length of a slice can grow and shrink as you see fit.
// Syntax: slice_name := []datatype{values
// len() function » returns the length of the slice (nº of elements in it)
// cap() function » returns the capacity of the slice (nº of elements it can grow or shrink to)

func main() {
	myslice1 := []int{}
	myslice2 := []string{"Hello", "Slices"}
	myslice3 := [3]int{6, 2: 10}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(len(myslice3))
	fmt.Println(cap(myslice3))
	fmt.Println(myslice3, n)

	// Create a Slice from an Array
	// var myarray = [legnth]datatype{values}
	// myslice := myarray[start:end]

	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	slic1 := arr1[2:4]

	fmt.Printf("slic1 = %v\n", slic1)
	fmt.Printf("length = %d\n", len(slic1))
	fmt.Printf("capacity = %d\n\n", cap(slic1)) // the slice starts, in this e.g., from the second element of an array that contains six elements, the slice can grow to the end of the array.

	slic1 = arr1[1:3]                     // Change length by re-slicing the array
	slic1 = append(slic1, 20, 21, 22, 23) // Change length by appending items

	// make() » Function to create a Slice

	myslice4 := make([]int, 5, 10) //syntex: slice_name := make([]type, length, capacity)
	myslice5 := make([]int, 5)     // if capacity is not defined, by default it is = length

	fmt.Printf("myslice4 = %v\n", myslice4)
	fmt.Printf("length = %d\n", len(myslice4))
	fmt.Printf("capacity = %d\n\n", cap(myslice4))

	fmt.Printf("myslice5 = %v\n", myslice5)
	fmt.Printf("length = %d\n", len(myslice5))
	fmt.Printf("capacity = %d\n\n", cap(myslice5))

	prices := []int{10, 20, 30}
	prices[2] = 50            // changes the third element from 20 to 50
	fmt.Println(prices[0])    // can access a specific slice element, just as in the array
	fmt.Println(prices[2], n) // in GO, indexes allways start at 0

	// Append Elements » if the slice has capacity it will add the elements to the end of the slices

	myslice6 := []int{1, 2, 3, 4, 5}
	fmt.Printf("myslice6 = %v\n", myslice6)
	fmt.Printf("length = %d\n", len(myslice6))
	fmt.Printf("capacity = %d\n\n", cap(myslice6))

	myslice6 = append(myslice6, 20, 21)
	fmt.Printf("myslice6 = %v\n", myslice6)
	fmt.Printf("length = %d\n", len(myslice6))
	fmt.Printf("capacity = %d\n\n", cap(myslice6))

	// Append Slices
	// Syntax: slice3 = append(slice1, slice2...) » the "..." is necessary

	myslice7 := []int{1, 2}
	myslice8 := []int{3, 4}
	myslice9 := append(myslice7, myslice8...) // Change length by appending items

	fmt.Printf("myslice9 = %v\n", myslice9)
	fmt.Printf("length = %d\n", len(myslice9))
	fmt.Printf("capacity = %d\n\n", cap(myslice9))

	//Memory Efficiency
	// copy() Function
	// Syntax = copy(dest, src) » copies data from src to dest. It returns the number of elements copied.

	numbers := []int{11, 12, 13, 14, 15} // original slice
	fmt.Printf("numbers = %v\n", numbers)

	// Copy only with only what is needed
	neededNumbers := numbers[:len(numbers)-2]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers) //copy function

	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n\n", cap(numbersCopy))

}
