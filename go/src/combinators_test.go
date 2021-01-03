package lambdacalculus

import (
	"testing"
)

func TestFlip_with_Pow(t *testing.T) {
	res := Flip(Pow)(four)(three).(xl)(f)(x)
	if res != 64 {
		t.Errorf("Power of four to three should be 64 instead is %v", res)
	}
}

func TestConstant(t *testing.T) {
	c := Constant(1)
	res := c(Succ(Zero))
	if res != 1 {
		t.Errorf("The constant should be always 1 instead is %v", res)
	}
}

// Test recursion obtained via Z combinator using a function which does use normal logic and not Lambda Calculus logic
func TestZWithFactorial(t *testing.T) {
	var factorial = func(f interface{}) interface{} {
		return func(n interface{}) interface{} {
			nInt := n.(int)
			if n == 0 {
				return 1
			}
			return nInt * f.(Lambda)(nInt-1).(int)
		}
	}
	fZRecursive := Z(factorial)
	res := fZRecursive.(Lambda)(5)
	if res != 120 {
		t.Errorf("The factorial of 5 should be always 120 instead is %v", res)
	}
}

// Test recursion obtained via Z combinator using a function which does use Lambda Calculus logic
func TestZWithFactorialLambda(t *testing.T) {
	var factorial = func(f interface{}) interface{} {
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
			var ifIsZero = func(n ChurchNumber) ChurchNumber {
				return one
			}
			// this is the function returned by IsZero if n is NOT zero
			var ifIsNOTZero = func(n ChurchNumber) ChurchNumber {
				fPrev := f.(Lambda)(Prev(nChurchNumber)).(ChurchNumber)
				return Mult(fPrev)(nChurchNumber)
			}
			return IsZero(nChurchNumber)(ifIsZero)(ifIsNOTZero).(nm)(nChurchNumber)
		}
	}
	// we pass the function to the Z combinator to obtain the factorial Lambda version
	fZRecursive := Z(factorial)
	// eventually we call the recursive Lambda function with the number we want to calculate the factorial for
	resChurch := fZRecursive.(Lambda)(four)
	res := resChurch.(ChurchNumber)(f)(x)
	if res != 24 {
		t.Errorf("The factorial of 4 should be always 24 instead is %v", res)
	}
}
