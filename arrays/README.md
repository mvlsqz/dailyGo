# Arrays
Arrays in Go are useful when planning for detailed layout of memory. Using arrays can sometimes help avoid allocation. However, their primary use is for the building blocks of slices.

## Initializing An Array
The capacity of an array is defined at creation time, and it can no
longer be changed.

## Array zero value
The zero value of each element in array is the zero value for the type
of elements in the array.

```go
name := [4]string{}
// ["" "" "" ""]
```

## Two Dimensional Arrays
```go
type Matrix [3][3]int

func main() {
  m := Matrix{
    {0, 0, 0},
    {1, 1, 1},
    {2, 2, 2},
  }
}
```

## For Loop
`for` loop is the only looping construct in Go

```
for i := 0; i < N; i++ {
  // do work until i equals N
}
```

## Continuing A Loop
The `continue` keyword allows us to go back to start of the loop and
stop executing the rest in the `for` block.

```go
for {
  if i == 3 {
    // go to the start of the loop
    continue
    // the current loop finished
  }

  // do work
}
```

## Breaking a Loop
To stop execution of a loop we can use the `break` keyword
```go
for {
  if i == 3 {
    // stop looping
    break
    // this will no longer run
  }

  // do work
}
```

## Do while loop
A `do while` loop is used in situations where at least the iteration
wanted, lets use for to simulate it.
```go
var i int
for {
  fmt.Println(i)
  i += 2
  if i >= 3 {
    break
  }
}
```

## `range` Keyword
```go
package main

import "fmt"

func main() {
  names := [4]string{"John", "Paul", "George", "Ringo"}

  for i, n := range names {
    fmt.Printf("%d - %s\n", i, n)
  }
}
```
