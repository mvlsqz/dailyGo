# Packages
Packages are how to Go organizes code. It is also how scope and
visibility are determined.

## What are Packages
A package represents all the files in a single directory on disk.
One directory can containe only files from the same package

## Package Names
Altough not required, strongly encouraged that package name to match the folder it is in.

**Source files in package must declare the package name at the top
of the file as the first code statement.**

## Executable Packages
Executable programs must have a `main` package that declares
a main() function

**The `main` function can only be declared once and receives no
arguments, nor does it return any values.**

## Package Resolution
For code that lives in local directory, lets say `blue/red`,
package name must be `red` and resolution should be
```
import "blue/red"
```

For code that lives on source code repositories, resolution must be:
```
import "github.com/gobuffalo/buffalo"
```

The source code would be in the following location on disk:
```
'${GOPATH}/src/github.com/gobuffalo/buffalo'
```

## Scope and Visibility
External visibility is controlled by capitalization. Types,
Variables, Functions, etc ... that start with a capital letter
are available, publicity, outside the current package.

A symbol that is visible outside its package is "exported"

Types, Variables, Functions, etc ... that start with a lower
case letter are unexported and are not available outside of current
package.

*All variables and types declared inside a package are visible
to everything else in the same package.*

## Security
It is important to understand that although you can't directly
access or change unexported fields in a struct, you can still
get access to view the contents of them.

This is because the `fmt` package make use of the
`reflection` package.

[example](./reflection/main.go)

[about Standard Package
Layout](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1)


## Understanding Circular Dependencies
A circular dependencie occurs when one package imports another package
that directly or indirectly imports the first package.

## Recomended Approach
One of the most common approaches to package layout is to separate out your application into 4 package types:
  - Domain Package
  - Implementation Packages
  - Mock Package
  - Binary Packages

[Recomended
read](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html)

### Domain Package
This package includes concrete data types and it includes interfaces for
services that your application relies on. It also includes functionality
that only relies on these data types and services

### Implementation Packages
Each implementation should isolate the technology inside of it. The package name is typically the technology used in the implementation. This separation of implementations makes it easy to inject mocks at each layer so you can unit test each layer separately.

### Mock Package
A single mock package allows all your packages to share mock
implementations of your domain interfaces. By using domain interfaces we
can test each layer of our application in isolation.

### Binary Package
A place to hold any generated binary for the application under `cmd`
package
