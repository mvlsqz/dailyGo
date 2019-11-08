package main

import "fmt"

func main() {
	type Movie struct {
		Title    string
		Released bool
		Length   int
	}

	movie := Movie{Title: "Wizard of OZ", Released: true, Length: 125}

	fmt.Println(movie)
}
