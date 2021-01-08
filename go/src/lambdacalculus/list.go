package lambdacalculus

// EmptyList is the empty list. The EmptyList responds true to all functions which are passed to it, even
// the function that is passed to check whether it is the empty list
var EmptyList List = func(f xl) interface{} {
	return T
}

// InsertOnTop inserts a value on top of a list
var InsertOnTop = func(x interface{}) func(l List) List {
	return func(l List) List {
		return func(f xl) interface{} {
			return f(x)(l)
		}
	}
}

// IsEmpty returns true if the list is empty, false otherwise
var IsEmpty = func(l List) ChurchBoolean {
	var alwayFalseUnlessListIsEmpty xl = func(x interface{}) Lambda {
		return func(y interface{}) interface{} {
			return F
		}
	}
	return l(alwayFalseUnlessListIsEmpty).(ChurchBoolean)
}

// ListElementVal returns the value of a list element. The value of the empty list is, by this convention, the empty list
var ListElementVal = func(element List) interface{} {
	return IsEmpty(element)(element)(element(First))
}

// NextListElement returns the next element in a list, if there is an element, otherwise return the empty list.
var NextListElement = func(element List) List {
	return IsEmpty(element)(element)(element(Second)).(List)
}
