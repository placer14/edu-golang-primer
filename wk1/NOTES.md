# Week 1 - Instructor Notes

## Overview

- Types in Go
- Pointers and value handling
- Advanced/complex types in Go
- Manipulating types/values

## Who reviewed the Go Tour ahead of class?

## Go Primitive Types

Intro about creating and manipulating variables:

- Declaring variables
- Zero values for each type

Walk through all of the primative types and their zero value.

- bool
- int (int8, int16, int32, int64)
- uint (uint8, uint16, uint32, uint64, uintptr)
- int, uint sizes (matches bitness of host arch)
- byte (uint8)
- rune (int32)
- float (float32, float64)
- complex (complex64, complex128)
- strings
  - literals ('hi', "hi", `hi`)
- Pointers
- Pass-by-value (pointers vs values)

## Go Composite Types

Walk through all of the complex types and their zero values:

- Slices
  - Indices (0-based)
  - Slicing slices. (half-open range, `[N:M]`, `[:M]`, `[N:]`, `[:]`)
  - string ([]byte, []rune)
  - Slices of ANYTHING (other slices!?)
- Arrays
  - Slicing
- Maps
- Functions
- Structs and methods (structural typing)
- Interfaces
  - `fmt.Stringer`
  - `error`
- Pointers of literals (complex works but not primatives)

## Language Keywords

- Just 25 keywords
  - Declaration: `const`, `func`, `import`, `package`, `type`, `var`
  - Unique language types notation: `chan`, `interface`, `map`, `struct`
  - Logic flow control: `break`, `case`, `continue`, `default`, `else`, `fallthrough`, `for`, `goto`, `if`, `range`, `return`, `select`, `switch`
  - Lifecycle flow control: `go`, `defer`
- Other built-ins which are not keywords: `panic`, `recover`, `len`, `cap`, `append`, `make`, `new`
- Control flow
  - Traditional: `if/else`, `for`, `switch`, `select`, etc
  - `panic` and `recover`
- Indentifier
  - An identifier is a token which must be composed of Unicode letters, Unicode digits and _(underscore), and start with either an Unicode letter or_.
  - Blank: `_`
  - Private vs Public (exported) identifier
- Operators (`+/-*`, `**`, `++`, `--`, `>`, `<`, `<=`, `>=`, `>>`, `<<`, `()`)
- Variable manipulation
  - new
  - make
  - for range
  - value mutation
  - Casting types
    - Overflow, underflow

## Missing Anything? Improvements?

## Share Contact Info
