package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 6; i++ {
		if i == 4 {
			continue
		}

		fmt.Println(10)
	}
}
