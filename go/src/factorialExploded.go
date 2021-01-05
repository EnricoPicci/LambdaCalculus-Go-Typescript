package lambdacalculus

// FactorialExpanded is the same function as Factorial but with all variables expanded to their minimal
// elements.
// This is to show that the factorial function can be obtained by pure composition of anonymous functions
var FactorialExpanded Lambda = func(g Lambda) Lambda {
	return func(x interface{}) interface{} {
		return g(func(v interface{}) interface{} {
			return (x.(Lambda)(x)).(Lambda)(v)
		})
	}(func(x interface{}) interface{} {
		return g(func(v interface{}) interface{} {
			return (x.(Lambda)(x)).(Lambda)(v)
		})
	}).(Lambda)
}(func(f interface{}) interface{} {
	return func(n interface{}) interface{} {
		return func(n ChurchNumber) ChurchBoolean {
			return n(func(x interface{}) Lambda {
				return func(y interface{}) interface{} {
					return x
				}
			}(func(x interface{}) Lambda {
				return func(y interface{}) interface{} {
					return y
				}
			}))(func(x interface{}) Lambda {
				return func(y interface{}) interface{} {
					return x
				}
			}).(ChurchBoolean)
		}(n.(ChurchNumber))(func(n ChurchNumber) ChurchNumber {
			return (func(n ChurchNumber) ChurchNumber {
				return func(f interface{}) Lambda {
					return func(x interface{}) interface{} {
						return f.(Lambda)(n(f)(x))
					}
				}
			})(func(f interface{}) Lambda {
				return func(x interface{}) interface{} {
					return x
				}
			})
		})(func(m ChurchNumber) ChurchNumber {
			return (func(n ChurchNumber) nm {
				return func(m ChurchNumber) ChurchNumber {
					return func(f interface{}) Lambda {
						return n(m(f))
					}
				}
			})(f.(Lambda)((func(n ChurchNumber) ChurchNumber {
				return n(func(p Pair) Pair {
					return (func(x interface{}) func(x interface{}) Pair {
						return func(y interface{}) Pair {
							return func(f xl) interface{} {
								return f(x)(y)
							}
						}
					})(p(func(x interface{}) Lambda {
						return func(y interface{}) interface{} {
							return y
						}
					}))((func(n ChurchNumber) ChurchNumber {
						return func(f interface{}) Lambda {
							return func(x interface{}) interface{} {
								switch f := f.(type) {
								case func(Pair) Pair:
									xp := x.(Pair)
									return f(n(f)(xp).(Pair))
								default:
									return f.(Lambda)(n(f)(x))
								}
							}
						}
					})(p(func(x interface{}) Lambda {
						return func(y interface{}) interface{} {
							return y
						}
					}).(ChurchNumber)))
				})((func(x interface{}) func(x interface{}) Pair {
					return func(y interface{}) Pair {
						return func(f xl) interface{} {
							return f(x)(y)
						}
					}
				})(func(f interface{}) Lambda {
					return func(x interface{}) interface{} {
						return x
					}
				})(func(f interface{}) Lambda {
					return func(x interface{}) interface{} {
						return x
					}
				})).(Pair)(func(x interface{}) Lambda {
					return func(y interface{}) interface{} {
						return x
					}
				}).(ChurchNumber)
			})(n.(ChurchNumber))).(ChurchNumber))(n.(ChurchNumber))
		}).(nm)(n.(ChurchNumber))
	}
})
