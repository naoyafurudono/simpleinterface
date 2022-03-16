# Memo

Note to study the language spec and how to implement the tool.

## Language spec

- [Syntax](https://go.dev/ref/spec#Type_declarations) of type declaration
- [Syntax](https://go.dev/ref/spec#Interface_types) of interface types

## libraries

- [go](https://pkg.go.dev/go@go1.18)
  - [ast](https://pkg.go.dev/go/ast)
  - [parser](https://pkg.go.dev/go/parser@go1.18)
    - use `ParserFile`
  - [type](https://pkg.go.dev/go/types@go1.18)
    - [interface](https://pkg.go.dev/go/types@go1.18#Interface)
- golang-set
  - [github](https://github.com/deckarep/golang-set)
  - [doc](https://pkg.go.dev/github.com/deckarep/golang-set)

>scienceSlice := []interface{}{"Biology", "Chemistry"}
>scienceClasses := mapset.NewSetFromSlice(scienceSlice)

### what want to do

1. extract interface definition from source code
2. extract type-set
