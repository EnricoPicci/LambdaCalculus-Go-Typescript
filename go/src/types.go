package lambdacalculus

// Lambda is a lambda function
type Lambda = func(x interface{}) interface{}

// xl is a function which takes a generic x and returns a Lambda
type xl = func(x interface{}) Lambda

// type depending on specific types defined in Lambda calculus (apart Lambda)
type nm = func(n ChurchNumber) ChurchNumber
type nxl func(n ChurchNumber) xl
type bb func(p ChurchBoolean) ChurchBoolean
