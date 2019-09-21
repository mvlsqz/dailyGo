# dailyGo
Daily Go Course

#Syntax and Types

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

`i := 1` use it when don't want Go go infer your data type, but want
use `short` method

**long form is considered not idiomatic** `var i int = 1`
