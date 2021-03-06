# Memo

Note to study the language spec and how to implement the tool.

## Language spec

- [Syntax](https://go.dev/ref/spec#Type_declarations) of type declaration
- [Syntax](https://go.dev/ref/spec#Interface_types) of interface types

### type set calculation

|interface definition|type set|
|:~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~|:~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~|
| `{ int }`                        | `int`                                  |
| `{ ~int }`                       | all type whose underlying type is `int` |
| `{}`                             | all non-interface types |
| `{ x; y; z }`                    | intersection of type set of `x`, `y`, and `z` |
| interface element                |    |
| `Method() tp`                    | types whose method set include the method |
| `X` where X is non-interface type | `{ X }`   |
| `~T`                              | `{ S | T is the underlying type of S}` |
| `t1|t2|...|tn`                    | union set of the type set of t1,...,tn |

## libraries

- [go](https://pkg.go.dev/go@go1.18)
  - [ast](https://pkg.go.dev/go/ast)
  - [parser](https://pkg.go.dev/go/parser@go1.18)
    - use `ParserFile`
  - [type](https://pkg.go.dev/go/types@go1.18)
    - [interface](https://pkg.go.dev/go/types@go1.18#Interface)
    - typeset : 欲しい機能がエクスポートされてない...
      - [code](https://cs.opensource.google/go/go/+/refs/tags/go1.18:src/go/types/typeset.go)
      - ライセンスがBSD likeで編集や再配布が許可されている。最悪コピペして取り込むか。どこまで取り込むか決めるのが大変そう。
      - どうやらtype set は取得しないでも、上手にAPIを駆使すればなんとかなりそう
- golang-set
  - NOTE this is not suitable for our need ;-(
  - use slice and write code by myself
  - [github](https://github.com/deckarep/golang-set)
  - [doc](https://pkg.go.dev/github.com/deckarep/golang-set)

### 静的解析ツールの作り方

- linterを作るなら skelton が便利（既に使っている）
  - go/analysis を使ったツールの雛形を作ってくれる
  - 抽象構文木と型データを自動で抽出する
  - テストとソースを書く
    - a以下のファイルをいじる
    - inspectorは深さ優先探索を再帰ではなくforで書けるようにいい感じにしてくれる
  - 型情報の取得には pass.TypesInfo.ObjectOfを使う

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

In general,  ai is an redundant element in union (a0 | ... | an) iff

```go
typeset( a0 |...| a(i-1) | a(i+1) |...| an ) = typeset( a0 |...| an )
```

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

In general,  ai is an redundant element in product (a0 ; ... ; an) iff

```go
typeset( a0 ;...; a(i-1) ; a(i+1) ;...;an ) = typeset( a0 ;...; an )
```

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
