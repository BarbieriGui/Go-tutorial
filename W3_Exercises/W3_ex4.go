package main

import (
	"fmt"
)

func main() {
	var cars [4]string = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	cars[0] = "Opel"
	fmt.Print(cars[0])
}

/*
to print an specific position use, e.g.: for the second position
	fmt.Print(cars[1])
	obs: array position ID start in 0,1,2..
and to change an position use, e.g.: to change "Volvo" to "Opel"
	cars[0]="Opel"
*/
