package main

import (
	"fmt"
)

func main() {
	n, _ := fmt.Println("this is a string", 4, true)

	// fmt.Println(a ...interface{})
	// Empty interface means give me as many arguments of any type (variadic argument)
	// fmt.Println use dot notation, it means from fmt i'll use Println
	fmt.Print(n)
}
