# Week 2 - Instructor Notes

## Overview

- Types in Go (continued)
- Pointers and value handling
- Advanced/complex types in Go
- Manipulating types/values

## Who reviewed the Go Tour ahead of class?

## Go Primitive Types

Intro about creating and manipulating variables:

- Declaring variables
- Zero values for each type

Walk through all of the primative types and their zero value.

- int (int8, int16, int32, int64)
- uint (uint8, uint16, uint32, uint64, uintptr)
- int, uint sizes (matches bitness of host arch)
- float (float32, float64)
- complex (complex32, complex64)
- byte (uint8)
- strings
  - literals ("hi", `hi`)
- rune (int32)
  - literals ('h')

--- we stopped at this point in wk1

- pickup some discussion about strings via <https://go101.org/article/string.html>

- bool
- complex (complex64, complex128)
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

--- we stopped at this point in wk2

- Channels
  - init: `make(chan`

## Language Keywords

- Just 25 keywords
  - Unique language types notation: `chan`, `interface`, `map`, `struct`
  - Logic flow control: `break`, `case`, `continue`, `default`, `else`, `fallthrough`, `for`, `goto`, `if`, `range`, `return`, `select`, `switch`
  - Declaration: `const`, `func`, `import`, `package`, `type`, `var`
  - Lifecycle flow control: `go`, `defer`
- Indentifier
  - An identifier is a token which must be composed of Unicode letters, Unicode digits and `_` (underscore), and start with either an Unicode letter or `_`.
  - Private vs Public (exported) identifier
- Control flow
  - Traditional: `if/else`, `for`, `switch`, `select`, etc
  - `panic` and `recover`

--- stop here for wk3

- Other built-ins which are not keywords: `panic`, `recover`, `len`, `cap`, `append`, `make`, `new`
- Operators (`+/-*`, `**`, `++`, `--`, `>`, `<`, `<=`, `>=`, `==`, `>>`, `<<`, `()`)
  - `==`
    - bool: regular equality
    - numeric (int, float, complex): regular equality
- Variable manipulation
  - new
  - make
  - for range
  - value mutation
  - Casting types
    - Overflow, underflow
  - Floats have approximation of arithmetic
    - Check for approximations when comparing Floats
    - Favor exp/root, then mul/div, before add/sub in a series of operations to improve accuracy
    - Group floating point operations into similar magnitudes where possible

- Convenience syntax
  - Blank: `_`
  - Numerical separator: `1_000_000`, `0b00_00_01`
  - Integer literals in other bases:
    - Octal (uses `o`) `0o017`
    - Binary (uses `b` or `B`): `0b011010`
    - Hex (uses `x` or `X`): `0xdeadbeef`
    - Imaginary (uses `i`): `i` == `sqrt(-1)`

## Missing Anything? Improvements?

## Share Contact Info
