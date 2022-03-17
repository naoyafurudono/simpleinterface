package a

func f() {
	type X interface {
		int | int64 | float64
	}

	type addable interface {
		int | int32 | int64 | float32 | float64
	}

	// type set is empty (danger!)
	type addable2 interface { // want "empty"
		int | int32 | int64 | float32
		float64
	} // -> {}

	type addables interface { // want "overwrap addable"
		addable
		addable
		addable
		addable
		addable
		addable
	} // --> addable = int64 | float64

	type Z interface { // want "overwrap X"
		addable
		X
		X
	} // --> int64

	type Myint int
	type x = interface{}
	type X1 interface {
		int64 | addable // ok
		~int
	} // --> {}

	type myStringer interface { // want "method and embedded type"
		String() string
		X1
	}

	type A interface {
		addable | X
	} // --> {int, int64, float64}

	type B1 interface {
		int
	}

	type B2 interface {
		int | int64
	}

	type B3 interface {
		B1 | B2 // = B2
	}

	// Myint or empty. MyintがString() string を実装しなかったらCの型セットは空
	// メソッドを実装するかは静的にわかることなので、この定義は怪しい
	// メソッドリストと~以外の basicではないインターフェース要素が同居しているときは警告
	type C interface { // want "method and embedded type"
		Myint
		String() string
	}

	// do not blame it!
	type Ok interface {
		~int
		String() string
	}

}
