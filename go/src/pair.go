package lambdacalculus

// Pair represents a pair
type Pair = fxlx

// Tuple2Struct is a 2-tuple structure
var Tuple2Struct = func(x interface{}) xp {
	return func(y interface{}) Pair {
		return func(f xl) interface{} {
			return f(x)(y)
		}
	}
}
