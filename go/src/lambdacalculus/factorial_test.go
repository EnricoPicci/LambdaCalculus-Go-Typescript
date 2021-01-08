package lambdacalculus

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	five := Succ(four)
	res := Factorial(five).(ChurchNumber)(f)(x)
	if res != 120 {
		t.Errorf("The factorial of 5 should be 120 instead is %v", res)
	}
}
