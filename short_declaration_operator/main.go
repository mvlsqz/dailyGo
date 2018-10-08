package main

import (
	"fmt"
)

func main() {
	// declaration and assignment
	x := 42
	fmt.Println(x)
	// just assignment because the variable already exists on line 9
	x = 99
	// we should have to use the variable :)
	fmt.Println(x)
	// A statement made up by an expression just for fun :P
	y := 100 + 24
	fmt.Println(y)
}
