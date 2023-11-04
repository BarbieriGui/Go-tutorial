package main

import (
	"fmt"
)

// GO Structure is used to store multiple values of different data types into a single variable.
/* Syntax
type struct_name struct {
	member1 datatype;
	member2 datatype;
	...
} */

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	var pers1 Person
	var pers2 Person

	// Pers1 specification
	pers1.name = "Marcos" // use (.) between structure variable name and member to access it
	pers1.age = 40
	pers1.job = "Teacher"
	pers1.salary = 5000

	// Pers2 specification
	pers2.name = "Rui"
	pers2.age = 30
	pers2.job = "Design"
	pers2.salary = 4800

	/*	Simple print option

		fmt.Println("Name: ", pers1.name)
		fmt.Println("Age: ", pers1.age)
		fmt.Println("Job: ", pers1.job)
		fmt.Println("Salary: ", pers1.salary)
		fmt.Println()

		fmt.Println("Name: ", pers2.name)
		fmt.Println("Age: ", pers2.age)
		fmt.Println("Job: ", pers2.job)
		fmt.Println("Salary: ", pers2.salary)*/

	// Print passing struc as a func argument
	printPerson(pers1)
	fmt.Println()
	printPerson(pers2)
}

func printPerson(pers Person) {
	fmt.Println("Name: ", pers.name)
	fmt.Println("Age: ", pers.age)
	fmt.Println("Job: ", pers.job)
	fmt.Println("Salary: ", pers.salary)
}
