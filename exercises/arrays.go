package main

import (
	"fmt"
)

func main() {
	fruits := [4]string{"Banana", "Orange", "Pineaple", "Strawberry"}
	for i := 0; i < len(fruits); i++ {
		fmt.Println(i, fruits[i])
	}

	for i, f := range fruits {
		println(i, f)
	}

	for i, f := range fruits {
		if i%2 == 0 {
			continue
		}

		fmt.Println(i, f)
	}
}
