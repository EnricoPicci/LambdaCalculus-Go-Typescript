package lambdacalculus

// T is boolean true
var T ChurchBoolean = First

// F is boolean false
var F ChurchBoolean = Second

// Not is the boolean not
var Not = func(p ChurchBoolean) xl {
	return Flip(p)
}

// And is the boolean and
var And = func(p ChurchBoolean) bb {
	return func(q ChurchBoolean) ChurchBoolean {
		return p(q)(p).(ChurchBoolean)
	}
}

// Or is the boolean and
var Or = func(p ChurchBoolean) bb {
	return func(q ChurchBoolean) ChurchBoolean {
		return p(p)(q).(ChurchBoolean)
	}
}
