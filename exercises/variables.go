package main

import "fmt"

func main() {
	var command string
	var valid bool

	foo := int(5)
	bar := bool(true)

	fmt.Printf("value of variable `command` is %s\n", command)
	fmt.Printf("value of variable `valid` is %v\n", valid)

	fmt.Printf("value of variable `foo` is: %d\n", foo)
	fmt.Printf("value of variable `bar` is: %v\n", bar)
}
