package main

import "fmt"

func main() {
	a := 1

	fmt.Println("This will join all arguments and print them. a = ", a)
	fmt.Printf("this is the `format` string. Scape \" and print of a: %v\n", a)

	type User struct {
		Name string
	}

	u := User{Name: "Mary"}

	fmt.Printf("user: %+v", u)
}
