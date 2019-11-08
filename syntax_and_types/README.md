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
```
`Email` starts with a uppercase letter and can be accessed by other
packages.

`password` starts with a lowercase letter, and is only
accesible inside the package it is declared in.

# Naming Style
* Use ver terse (or short) variable names

* The smaller the scope the variable exists in, the smaller the variable name
```go
names := []string{"Mary", "John", "Bob", "Anna"}
for i, n := range names {
  fmt.Printf("index: %d = %q\n", i, n)
}
```
`names` is used in a larger scope, so it would be common to give it a more meaningful name to help remember what means in the program.

`i`, `n` are used inmediately in the next line of code, and never used again

### Notes about style
- Underscores aren't conventional
- Short over larger variable names
- acronyms should be capitalized

# Multiple Assignment
Go allows assign several values to several variables within same line.

`j, k, l :=  "gopher", 2.05, 15`

**Note** This approach can keep your code lines down, but make sure you're no compromising redability for fewer lines of code.

# Statically Typed
Go is statically typed language. Statically typed means that each statement in the program is checked at compile time. It also means that the data type is bound to the variable, whereas in Dynamically linked languages, the data type is bound to the value.

# Numeric Type
## architecture independent type
`uint8` `uint16` `uint32` `uint64` `int8` `int16` `int32` `int64` `float32` `float64` `complex64` `complex128` `byte` `rune`

This types will have the correct size regardless of the architecture that the progran is compiled for

## implementation specific type
`uint` `int` `uintptr`
This types will have their size defined by the architecture the program is compiled for

# Booleans
It's useful for making logic decisions in code

# Strings
Sequence of one or more characters (letters, numbers, symbols), that can
be either a constant or variable, i.e.
```go
`raw string literal` // note back quotes
"interpreted string litteral" // double quotes
```
## Raw String Literal
Raw string literals are character sequences between back quotes, often
called back ticks. Within the quotes, any character may appear except
back quote. Backslash doesn't have any special meaning inside back
quote. Raw String literal may also be used to create multilne strings

## Interpreted String Literal
Will almost always use interpreted string literals because they allow for escape characters within them.

# UTF8
Go builtin support for UTF8, i.e:
`a := "Hello, 世界"`

Notice that UTF8 have some points to be aware of:
```go
a := "Hello, 世界"
for i, c := range a {
	fmt.Printf("%d: %s\n", i, string(c))
}
fmt.Println("length of 'Hello, 世界': ", len(a))
```
Output
```bash
0: H
1: e
2: l
3: l
4: o
5: ,
6:
7: 世
10: 界
length of 'Hello, 世界':  13
```
See The length is longer than the number of times it took to range over the string. The reason for this is because runes in Go are a special type. A rune is an alias for int32. A rune can be made up of 1 to 3 int32 values.

# Constants
Constants are like variables, except they can't be modified once they
have been declared
## Untyped
Constants can be untyped. This can be useful when working with numbers such as integer type data. If the constant is untyped, it is explicitly converted, where typed constants are not.

## Typed
In the other hand typed constants can only operates with data types that
are the same as it's data type assigned on declaration time.

# IOTA
Simplify definitions of incrementing numbers and is used on `const`
block. Because it can be used in expressions, it provides a generality
beyond that of simple enumerations.

**For `iota` order matters**

# Verbs
Verbs that are used with `Print` will create different style of
formatting

```bash
// Use %v to print the value
// Use %s to print a string
// Use %q to quote a string
// Use %d to print a decimal (int, int32, etc)
// Use %T to print out the data type of the variable
// Use \ to escape a character, like a quote:  \"
// Use \n to print a new line (line return)
// Use \t to insert a tab
// Use %+v to print the name and value
```
# Structs

A struct is a collection of fields, often called members (or attributes).

Structs are used to create custom complex types in Go.

When trying to understand structs, it helps to think of them as a blueprint for what the new type will do.

A struct definition, does NOT contain any data.

**[STYLE GUIDE](https://github.com/uber-go/guide/blob/master/style.md)**
