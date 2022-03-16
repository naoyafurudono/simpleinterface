package a

func f() {
	type X interface {
		int | int64 | float64
	}
	type addable interface { // want "overwrap int64" "overwrap float64"
		int | int32 | int64 | float32 | float64
		float64 | int64
	}
	type Y interface { // want "overwrap addable"
		addable
		addable
		addable
		addable
		addable
		addable
	}
	type Z interface {
		addable
		X
	}

	type X1 interface {
		int
		~int
	}
}
