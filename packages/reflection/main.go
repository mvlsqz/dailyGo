package main

import (
	"fmt"

	"github.com/mvlsqz/dailyGo/packages/reflection/foo"
)

func main() {
	user := foo.NewUser("Cory", "LaNou", "password")
	fmt.Printf("%+v\n", user)
	// can't be accessed directly
	// fmt.Print(user.password)
	// can't be modified directly
	// user.password = "newValue"
	// Output: user.password undefined (type foo.user has no field or method password)
}
