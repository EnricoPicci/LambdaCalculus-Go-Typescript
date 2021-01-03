package lambdacalculus

// Lambda is a lambda function
type Lambda = func(x interface{}) interface{}

// type names are built using the following convetion - consider that we are typing "unary functions" since
// this is what lambda functions are
//
// x denotes a generic interface
// l denotes a Lambda as defined above (a unary function that accepts a generic input and retruns a generic output)
// f is used to denote a parameter as a function
// Types can be used to compose other types, so xlx denotes a function that takes an x and returns an lx
type xl func(x interface{}) Lambda
type fxlx func(f xl) interface{}

// type depending on specific types defined in Lambda calculus (apart Lambda)
type nm func(n ChurchNumber) ChurchNumber
type nxl func(n ChurchNumber) xl
type bb func(p ChurchBoolean) ChurchBoolean
type xp func(x interface{}) Pair
