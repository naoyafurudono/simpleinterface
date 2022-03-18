package a

func f() {
	type basic interface {
		Method1() string
		Method2()
	}

	type X interface {
		int | int64 | float64
	}

	type addable interface {
		int | int32 | int64 | float32 | float64
	}

	// type set is empty (danger!)
	type empty interface { // TODOwant "empty"
		int | int32 | int64 | float32
		float64
	}

	// --------- over wrap ---------

	// do not blame many times
	type addables interface { // want "overwrap addable"
		addable
		addable
		addable
		addable
		addable
		addable
	}

	type Z interface { // want "overwrap X"
		addable
		X
		X
	} // --> int64

	// --------- method with no method ---------

	// addable or empty. addableがString() string を実装しなかったら型セットは空
	// このインターフェースを実装する型はaddableの型セットに属していて、String() string を持たないといけない
	// addableの型セットはコンパイル時に決まるので、このインターフェースは不自然
	// メソッドリストと~以外の basicではないインターフェース要素が同居しているときは警告
	type myStringer interface { // want "method and embedded type"
		String() string
		addable
	}

	type ok interface { // ok
		String() string
		basic
	}

	// do not blame it!
	type Ok interface {
		~int
		String() string
	}

	type intlike interface {
		~int
	}

	type corner interface { // ok
		intlike
		String() string
	}
}

func g() {
	type MyInt interface {
		int
	}

	type MyInt2 interface {
		int
	}

	type MyMyInt interface {
		MyInt
		~int
	} // MyInt

}
