# Lambda Calculus implemented with Go and Typescript

This project contains some Lambda Calculus computations implemented both in Go and in Typescript.

- Natural Numbers as per Chrch encoding with some arithmetics: Sum, Mult, Pow, Sub
- Boolean as per Chrch encoding and boolean logic operators: And, Or, Not
- Recursion via Z-fixed-point combinator through which Factorial calculation is implemented

## Go

`lambdacalculus` package contains all the code split over different files.

- [types.go](./src/lambdacalculus/types.go): all types and type aliases defined here
- [church-booleans.go](./src/lambdacalculus/church-booleans.go): boolean operators for Church encoded booleans
- [churhc-numbers.go](./src/lambdacalculus/church-numbers.go): arithmetic funcions for Church encoded natural numbers
- [pair.go](./src/lambdacalculus/pair.go): 2-tuple (Pair) structure
- [list.go](./src/lambdacalculus/list.go): operations on lists
- [combinators.go](./src/lambdacalculus/combinators.go): some combinators including the Z combinator
- [factorial.go](./src/lambdacalculus/factorial.go): factorial function

`lambdacalculusexpanded` package contains the implementation of the factorial function using only anonymous functions (function literals) and generic interface{} types

### Test

- Go to the Go folder with `cd go`
- Launch the tests with the command `go test ./...`

### Interactive testing and debugging from within VSCode

- Open the test file with VSCode
- Open the "Run" view, i.e. the VSCode view which allows to run programs
- From the top lesft drop-down list, select "GO Launch"
- If required, breakpoints can be placed in the open file in which case the execution stops at the breakpoints and allows interactive
  debugging

## Typescript

### Test

- Go to the Go folder with `cd typescript`
- Launch the tests with the command `npm test`

### Interactive testing and debugging from within VSCode

- Open the test file with VSCode
- Open the "Run" view, i.e. the VSCode view which allows to run programs
- From the top lesft drop-down list, select "Current TS Test File"
- If required, breakpoints can be placed in the open file in which case the execution stops at the breakpoints and allows interactive
  debugging
