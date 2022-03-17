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

You can see the current ability at the test code.
<!-- TODO testへのリンク-->

## Current status

Implement nothing.

## Plan to implement

### Existing analysis by the compiler

Some kind of redundant definitions are detected by the Go compiler. They are out of our scope now.

```go
type Done interface {
    addable | addable // compiler blames it
}
```

### DONE Simple redundancy

Following is the simplest target example of this tool.

```go
type Xs interface {
        addable
        addable
}
```

### Complex case

How about following.

```go
// 再掲
type addable interface {
        int | int32 | int64 | float32 | float64 | string
}

// 再掲
type mulable interface {
        int | int32 | int64 | float32 | float64
}

type integer interface {
        int | int32 | int64
}

type float interface {
        float32 | float64
}

type number interface {
        integer | float
}

// same as addable (is it useful to be detected?)
type nums interface {
    number | string
}
```

## Discussion about implementation

see `memo.md`
