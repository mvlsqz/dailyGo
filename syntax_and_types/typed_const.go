package main

import "fmt"

const (
	leapYear = int32(366) // typed
)

func main() {
	hours := 24
	fmt.Println(hours * leapYear) // multiplying int and int32 types
	// this will return `mismatched types int and int32` error
}
