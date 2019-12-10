# Optional

This is a collection of optional wrappers for primitive Go types and the code generator of optional wrappers for arbitraty types.

## Usage

To generate optional wrappers for your types using `go generate` utility, put a comment `//go:generate go run github.com/eqld/optional/cmd <type1> <type2> ...` at the beggining of any file in a package, or annotiate arbitrary global type definition started from `type` keyword in your code. You can combine both approaches. Generated files will be written into current working directory ("PWD" environment variable).
