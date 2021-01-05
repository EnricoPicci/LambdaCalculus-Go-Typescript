package lambdacalculus

// Pair represents a pair
type Pair func(f xl) interface{}

// Tuple2Struct is a 2-tuple structure
var Tuple2Struct = func(x interface{}) func(x interface{}) Pair {
	return func(y interface{}) Pair {
		return func(f xl) interface{} {
			return f(x)(y)
		}
	}
}
