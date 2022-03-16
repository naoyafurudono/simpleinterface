# Simple Interface

Simple interface is static analysis tool for Go.
The purpose of this tool is to warn complected definitions of interfere, which is
extended in Go1.18. For instance, following code can is legal Go program and compiled without any warning from the Go compiler.

```go
type addable interface {
        int | int32 | int64 | float32 | float64 | string
}

type mulable interface {
        int | int32 | int64 | float32 | float64
}

type X interface {
        addable | mulable
        addable
        mulable
}

func id[T X](elem T) T {
        return elem
}

func main() {
        id(1)
}
```

We can define `X` in more simple way:

```go
type X0 interface {
    addable
    mulable
}
```

Simple Interface will detect redundant interface definitions like of `X` and
possibly modify it to simpler one (like `X0`).

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

### Simple redundancy

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
