# dailyGo
Daily Go Course

# Syntax and Types

## reserved Keywords
`break` `default` `func` `interface` `case` `defer` `go` `map`
`chan` `else` `goto` `package` `const` `fallthrough` `if` `range`
`continue` `for` `import` `return`

[Go Documentation](https://golang.org/ref/spec#Keywords)

## Operators and delimiters
`+` `&` `+=` `&=` `&&` `==` `!=` `(` `)`
`-` `|` `-=` `|=` `||` `<` `<=` `[` `]`
`*` `^` `*=` `^=` `<-` `>` `>=` `{` `}`
`/` `<<` `/=` `<<=` `++` `=` `:=` `,` `;`
`%` `>>` `%=` `>>=` `--` `--` `!` `...` `.` `:`
`&^` `&^=`

[Go Documentation](https://golang.org/ref/spec#Operators_and_punctuation)
# Declaring Variables
there are several ways to declare a variable, and in some cases,
more than one way to declare the exact same variable and value.

`var i int` simple declaration

`var i int = 1` declaration and initialization

`i := 1` declaration using `short` method

### idiomatic way

`var i int` use only when initialization is not needed

`i := 1` use when need declaration and intialization

`i := int64(1)` use it when don't want Go go infer your data type, but want
use `short` method

**long form is considered not idiomatic** `var i int = 1`

# Zero values
All builtin types have a zero value. Any allocated variable is
usable even if it never has a value assigned.

In Go, because all values have a zero value, you can't have `undefined` values like some other languages.

For instance, a `boolean` in some languages could be `undefined`, `true`, or `false`, this allowing for three states to the variable. In Go, you can't have more than two states for a boolean value.

# Naming Rules
* Variable names must only be one word
* Variable names must be made up of only letters, numbers, and
  underscore
* Variable names cannot begin with a number

**variables are case-sensitive**, but is important to avoid using
similar variable names within a program

### First letter of a variable has special meaning in Go
- If variable starts with an uppercase letter, then that variable
  is accessible outside the package it was declared in.
- If variable starts with lowercase letter, then is only
  available within the package it is declared in.

```go
var Email string
var password string
``
`Email` starts with a uppercase letter and can be accessed by other
packages.
`password` starts with a lowercase letter, and is only
accesible inside the package it is declared in.
