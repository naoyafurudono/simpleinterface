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
    - typeset : 欲しい機能がエクスポートされてない...
      - [code](https://cs.opensource.google/go/go/+/refs/tags/go1.18:src/go/types/typeset.go)
- golang-set
  - NOTE this is not suitable for our need ;-(
  - use slice and write code by myself
  - [github](https://github.com/deckarep/golang-set)
  - [doc](https://pkg.go.dev/github.com/deckarep/golang-set)

### what want to do

- DONE extract interface definition from source code
- DONE extract type-set
- redundant element in a union
- redundant element in a sequence

#### redundant element in a union

```go
type X interface {
  int
}

type Y interface {
  int | int64
}

type Z interface {
  X | Y  // = Y
}
```

X is redundant in Z.

If we combine further, the problem become not easy to solve.

#### redundant element in a sequence

```go
type X interface {
  int
}

type Y interface {
  int | int64
}

type Z interface {
  X
  Y
} // the type set is same as X
```

Y is redundant in Z.

#### take care of underlying types

it may advanced?

```go
type X interface {
  int
}

type Y interface {
  int | int64
}

type Z interface {
  X
  Y
  ~int
} // the type set is same as X
```
