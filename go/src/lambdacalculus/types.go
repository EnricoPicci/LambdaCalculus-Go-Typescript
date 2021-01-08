package lambdacalculus

// Lambda is a lambda function
type Lambda = func(x interface{}) interface{}

// xl is a function which takes a generic x and returns a Lambda
type xl func(x interface{}) Lambda

// ChurchNumber are Church encoded numbers
type ChurchNumber func(f interface{}) Lambda

// ChurchBoolean are Church encoded booleans
type ChurchBoolean xl

// Pair represents a pair
type Pair func(f xl) interface{}

// List is a Pair whose first value is the value of the current element and the second value is the next element of the list,
// which is again a Pair
type List Pair

// type depending on specific types defined in Lambda calculus (apart Lambda)
type nm func(n ChurchNumber) ChurchNumber
type nxl func(n ChurchNumber) xl
type bb func(p ChurchBoolean) ChurchBoolean
