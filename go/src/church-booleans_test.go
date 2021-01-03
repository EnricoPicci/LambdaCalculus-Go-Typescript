package lambdacalculus

import (
	"testing"
)

func TestNotTrue(t *testing.T) {
	res := Not(T)(true)(false)
	if res.(bool) {
		t.Errorf("Not true should be false instead is %v", res)
	}
}

func TestNotFalse(t *testing.T) {
	res := Not(F)(true)(false)
	if !res.(bool) {
		t.Errorf("Not false should be true instead is %v", res)
	}
}

func TestAnd_T_T(t *testing.T) {
	res := And(T)(T)(true)(false)
	if !res.(bool) {
		t.Errorf("And(true)(true) should be true instead is %v", res)
	}
}
func TestAnd_T_F(t *testing.T) {
	res := And(T)(F)(true)(false)
	if res.(bool) {
		t.Errorf("And(true)(false) should be false instead is %v", res)
	}
}
func TestAnd_F_T(t *testing.T) {
	res := And(F)(T)(true)(false)
	if res.(bool) {
		t.Errorf("And(false)(true) should be false instead is %v", res)
	}
}
func TestAnd_F_F(t *testing.T) {
	res := And(F)(F)(true)(false)
	if res.(bool) {
		t.Errorf("And(false)(false) should be false instead is %v", res)
	}
}

func TestOr_T_T(t *testing.T) {
	res := Or(T)(T)(true)(false)
	if !res.(bool) {
		t.Errorf("Or(true)(true) should be true instead is %v", res)
	}
}
func TestOr_T_F(t *testing.T) {
	res := Or(T)(F)(true)(false)
	if !res.(bool) {
		t.Errorf("Or(true)(false) should be true instead is %v", res)
	}
}
func TestOr_F_T(t *testing.T) {
	res := Or(F)(T)(true)(false)
	if !res.(bool) {
		t.Errorf("Or(false)(true) should be true instead is %v", res)
	}
}
func TestOr_F_F(t *testing.T) {
	res := Or(F)(F)(true)(false)
	if res.(bool) {
		t.Errorf("Or(false)(false) should be false instead is %v", res)
	}
}
