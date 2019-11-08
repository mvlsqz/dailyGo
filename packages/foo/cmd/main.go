package main

import (
	"fmt"

	"github.com/mvlsqz/dailyGo/packages/foo"
)

func main() {
	user := foo.User{
		First: "Rob",
		Last:  "Pike",
	}

	fmt.Println(user)
}
