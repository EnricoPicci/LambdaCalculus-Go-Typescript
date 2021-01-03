package lambdacalculus

import (
	"testing"
)

func TestPair_First(t *testing.T) {
	res := Tuple2Struct(1)(2)(First)
	if res != 1 {
		t.Errorf("First of pair(1)(2) should be 1 instead is %v", res)
	}
}
func TestPair_Second(t *testing.T) {
	res := Tuple2Struct(1)(2)(Second)
	if res != 2 {
		t.Errorf("Second of pair(1)(2) should be 2 instead is %v", res)
	}
}

func TestPair_ChurchNumbers_First(t *testing.T) {
	res := Tuple2Struct(one)(two)(First).(ChurchNumber)(f)(x)
	if res != 1 {
		t.Errorf("First of pair(one)(two) should be 1 instead is %v", res)
	}
}
func TestPair_ChurchNumbers_Second(t *testing.T) {
	res := Tuple2Struct(one)(two)(Second).(ChurchNumber)(f)(x)
	if res != 2 {
		t.Errorf("Second of pair(one)(two) should be 2 instead is %v", res)
	}
}
