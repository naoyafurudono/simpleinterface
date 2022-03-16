package a

func f() {
	type X interface { // want "type spec"
		int
	}
	type addable interface { // want "type spec" "overwrap int64" "overwrap float64"
		int | int32 | int64 | float32 | float64
		float64 | int64
	}
}
