package lambdacalculus

import (
	"testing"
)

func TestIsEmptyWithEmpty(t *testing.T) {
	res := IsEmpty(EmptyList)(true)(false)
	if !res.(bool) {
		t.Errorf("The empty list should be empty instead is %v", res)
	}
}
func TestValueOfEmpty(t *testing.T) {
	val := ListElementVal(EmptyList)
	res := IsEmpty(val.(List))(true)(false)
	if !res.(bool) {
		t.Errorf("The value of empty list should be the empty list instead is %v", res)
	}
}
func TestNextOfEmpty(t *testing.T) {
	nextEl := NextListElement(EmptyList)
	res := IsEmpty(nextEl)(true)(false)
	if !res.(bool) {
		t.Errorf("The next element of empty list should be empty list instead is %v", res)
	}
}

func TestIsEmptyWithNotEmpty(t *testing.T) {
	list := InsertOnTop(one)(EmptyList)
	res := IsEmpty(list)(true)(false)
	if res.(bool) {
		t.Errorf("IsEmpty applied to a non empty list should be false instead is %v", res)
	}
}

func TestValueWithNotEmpty(t *testing.T) {
	list := InsertOnTop(one)(EmptyList)
	list = InsertOnTop(two)(list)
	val := ListElementVal(list).(ChurchNumber)
	res := val(f)(x)
	if res != 2 {
		t.Errorf("The current value of the list should be 2 instead is %v", res)
	}
}

func TestNextOfNotEmpty(t *testing.T) {
	list := InsertOnTop(one)(EmptyList)
	list = InsertOnTop(two)(list)
	nextEl := NextListElement(list)
	val := ListElementVal(nextEl).(ChurchNumber)
	res := val(f)(x)
	if res != 1 {
		t.Errorf("The value of the next elemenet should be 1 instead is %v", res)
	}
}

func TestLsatNextOfNotEmpty(t *testing.T) {
	// build a list with 3 elements: 3, 2 and 1
	list := InsertOnTop(three)(InsertOnTop(two)(InsertOnTop(one)(EmptyList)))
	// get the next element starting from 3 for three times: the first time you get 2, the second time you get 1 and the third time
	// you should get the empty list, which represents the end of the list
	nextEl := NextListElement(NextListElement(NextListElement(list)))
	res := IsEmpty(nextEl)(true)(false)
	if !res.(bool) {
		t.Errorf("IsEmpty applied to the end element of a list should be true instead is %v", res)
	}
}
