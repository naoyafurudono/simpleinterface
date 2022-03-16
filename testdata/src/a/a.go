package a

func f() {
	type X interface {
		int | int64 | float64
	} // --> int | int64 | float64

	type addable interface { // want "overwrap int64" "overwrap float64"
		int | int32 | int64 | float32 | float64
		float64 | int64
	} // int64 | float64

	// type set is empty (danger!)
	type addable2 interface {
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

	type myStringer interface {
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
}
