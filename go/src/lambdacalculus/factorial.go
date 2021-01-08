package lambdacalculus

var _factorial = func(f interface{}) interface{} {
	return func(n interface{}) interface{} {
		// the logic implemented to calculate the factorial is the following
		// if n is zero, then return one
		// if n is not zero, then return the multiplication of n * f(n-1)
		//
		// the implementation of this logic using Lambda Calculus is the following
		// IsZero(n)(one)(Mult(n)(f(Prev(n))))
		nChurchNumber := n.(ChurchNumber)
		//
		// to avoid infinite loop that comes because of the eager nature or Go lang, rather than passing "one" and "Mult(n)(f(Prev(n)))"
		// as parameters of "IsZer(n)" we must pass 2 functions like these "n -> one" and "n -> Mult(n)(f(Prev(n)))"
		// in this way we avoid Go trying to evaluate "Mult(n)(f(Prev(n)))" at each iteration, which is what causes the infinite loop
		// since it contains a reference to the "f" function which is passed in as parameter
		//
		// this is the function returned by IsZero if n is zero
		var ifIsZero nm = func(n ChurchNumber) ChurchNumber {
			return Succ(Zero)
		}
		// this is the function returned by IsZero if n is NOT zero
		var ifIsNOTZero nm = func(n ChurchNumber) ChurchNumber {
			fPrev := f.(Lambda)(Prev(nChurchNumber)).(ChurchNumber)
			return Mult(fPrev)(nChurchNumber)
		}
		return IsZero(nChurchNumber)(ifIsZero)(ifIsNOTZero).(nm)(nChurchNumber)
	}
}

// Factorial is the result of passing _factorial to the Z combinator to obtain the factorial Lambda function
var Factorial Lambda = Z(_factorial)
