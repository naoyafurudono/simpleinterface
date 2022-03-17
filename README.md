# Simple Interface

Simple interface is static analysis tool for Go1.18.
The purpose of this tool is to warn complected or unintended definitions of interfaces.
For instance, following code is legal Go program and compiled without any warning.

```go
type addable interface {
        int | int32 | int64 | float32 | float64 | string
}

type mulable interface {
        int | int32 | int64 | float32 | float64
}

// XXX the type set is same as addable...
type X interface {
        addable | mulable
        addable
}

func id[T X](elem T) T {
        return elem
}

func main() {
        id(1)
}
```

Simple Interface will detect redundant interface definitions like of `X` and
possibly (suggest?) modify it to simpler one (like `X0`).
The scope of this tool will be gradually extended.

## Installation

<!-- TODO -->

## Usage

<!-- TODO -->

### sample

You can see the current ability at the
[test code](https://github.com/naoyafurudono/simpleinterface/blob/main/testdata/src/a/a.go).

## Discussion about implementation

see [memo.md](https://github.com/naoyafurudono/simpleinterface/blob/main/memo.md).
