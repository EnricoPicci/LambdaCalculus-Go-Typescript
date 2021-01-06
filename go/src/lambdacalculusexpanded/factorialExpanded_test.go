package lambdacalculusexpanded

import (
	"testing"
)

// This tests shows that it is possible to define a function, in this case the factorial of a natural number,
// as well as a number, which in this test is the number 4, using only functions and therefore in pure
// Lambda Calculus terms

var incrementByOne = func(x interface{}) interface{} {
	return x.(int) + 1
}
var startValue = 0

// In this example, apart the above defined incrementByOne and startValue that are used to map Church encoded natural numbers
// to natural numbers as we know them, we end up using only anonimous functions (i.e. the literals func(...) {...} )
// and generic interfaces (i.e. interface{}). No other construct is used.

func TestFactorialExpanded(t *testing.T) {

	res := func(f func(x interface{}) interface{}) func(x interface{}) interface{} {
		return func(x interface{}) interface{} {
			return f(func(v interface{}) interface{} {
				return (x.(func(x interface{}) interface{})(x)).(func(x interface{}) interface{})(v)
			})
		}(func(x interface{}) interface{} {
			return f(func(v interface{}) interface{} {
				return (x.(func(x interface{}) interface{})(x)).(func(x interface{}) interface{})(v)
			})
		}).(func(x interface{}) interface{})
	}(func(f interface{}) interface{} {
		return func(n interface{}) interface{} {
			return func(n func(f interface{}) func(x interface{}) interface{}) func(x interface{}) func(x interface{}) interface{} {
				return n(func(x interface{}) func(x interface{}) interface{} {
					return func(y interface{}) interface{} {
						return x
					}
				}(func(x interface{}) func(x interface{}) interface{} {
					return func(y interface{}) interface{} {
						return y
					}
				}))(func(x interface{}) func(x interface{}) interface{} {
					return func(y interface{}) interface{} {
						return x
					}
				}).(func(x interface{}) func(x interface{}) interface{})
			}(n.(func(f interface{}) func(x interface{}) interface{}))(func(n func(f interface{}) func(x interface{}) interface{}) func(f interface{}) func(x interface{}) interface{} {
				return (func(n func(f interface{}) func(x interface{}) interface{}) func(f interface{}) func(x interface{}) interface{} {
					return func(f interface{}) func(x interface{}) interface{} {
						return func(x interface{}) interface{} {
							return f.(func(x interface{}) interface{})(n(f)(x))
						}
					}
				})(func(f interface{}) func(x interface{}) interface{} {
					return func(x interface{}) interface{} {
						return x
					}
				})
			})(func(m func(f interface{}) func(x interface{}) interface{}) func(f interface{}) func(x interface{}) interface{} {
				return (func(n func(f interface{}) func(x interface{}) interface{}) func(n func(f interface{}) func(x interface{}) interface{}) func(f interface{}) func(x interface{}) interface{} {
					return func(m func(f interface{}) func(x interface{}) interface{}) func(f interface{}) func(x interface{}) interface{} {
						return func(f interface{}) func(x interface{}) interface{} {
							return n(m(f))
						}
					}
				})(f.(func(x interface{}) interface{})((func(n func(f interface{}) func(x interface{}) interface{}) func(f interface{}) func(x interface{}) interface{} {
					return n(func(p func(f func(x interface{}) func(x interface{}) interface{}) interface{}) func(f func(x interface{}) func(x interface{}) interface{}) interface{} {
						return (func(x interface{}) func(x interface{}) func(f func(x interface{}) func(x interface{}) interface{}) interface{} {
							return func(y interface{}) func(f func(x interface{}) func(x interface{}) interface{}) interface{} {
								return func(f func(x interface{}) func(x interface{}) interface{}) interface{} {
									return f(x)(y)
								}
							}
						})(p(func(x interface{}) func(x interface{}) interface{} {
							return func(y interface{}) interface{} {
								return y
							}
						}))((func(n func(f interface{}) func(x interface{}) interface{}) func(f interface{}) func(x interface{}) interface{} {
							return func(f interface{}) func(x interface{}) interface{} {
								return func(x interface{}) interface{} {
									switch f := f.(type) {
									case func(func(f func(x interface{}) func(x interface{}) interface{}) interface{}) func(f func(x interface{}) func(x interface{}) interface{}) interface{}:
										return f(n(f)(x.(func(f func(x interface{}) func(x interface{}) interface{}) interface{})).(func(f func(x interface{}) func(x interface{}) interface{}) interface{}))
									default:
										return f.(func(x interface{}) interface{})(n(f)(x))
									}
								}
							}
						})(p(func(x interface{}) func(x interface{}) interface{} {
							return func(y interface{}) interface{} {
								return y
							}
						}).(func(f interface{}) func(x interface{}) interface{})))
					})((func(x interface{}) func(x interface{}) func(f func(x interface{}) func(x interface{}) interface{}) interface{} {
						return func(y interface{}) func(f func(x interface{}) func(x interface{}) interface{}) interface{} {
							return func(f func(x interface{}) func(x interface{}) interface{}) interface{} {
								return f(x)(y)
							}
						}
					})(func(f interface{}) func(x interface{}) interface{} {
						return func(x interface{}) interface{} {
							return x
						}
					})(func(f interface{}) func(x interface{}) interface{} {
						return func(x interface{}) interface{} {
							return x
						}
					})).(func(f func(x interface{}) func(x interface{}) interface{}) interface{})(func(x interface{}) func(x interface{}) interface{} {
						return func(y interface{}) interface{} {
							return x
						}
					}).(func(f interface{}) func(x interface{}) interface{})
				})(n.(func(f interface{}) func(x interface{}) interface{}))).(func(f interface{}) func(x interface{}) interface{}))(n.(func(f interface{}) func(x interface{}) interface{}))
			}).(func(n func(f interface{}) func(x interface{}) interface{}) func(f interface{}) func(x interface{}) interface{})(n.(func(f interface{}) func(x interface{}) interface{}))
		}
		// the lines below represent the number 4, as per Church encoding, which is the input provided to the factorial function
	})(func(n func(f interface{}) func(x interface{}) interface{}) func(f interface{}) func(x interface{}) interface{} {
		return func(f interface{}) func(x interface{}) interface{} {
			return func(x interface{}) interface{} {
				switch f := f.(type) {
				case func(func(f func(x interface{}) func(x interface{}) interface{}) interface{}) func(f func(x interface{}) func(x interface{}) interface{}) interface{}:
					return f(n(f)(x.(func(f func(x interface{}) func(x interface{}) interface{}) interface{})).(func(f func(x interface{}) func(x interface{}) interface{}) interface{}))
				default:
					return f.(func(x interface{}) interface{})(n(f)(x))
				}
			}
		}
	}(func(n func(f interface{}) func(x interface{}) interface{}) func(f interface{}) func(x interface{}) interface{} {
		return func(f interface{}) func(x interface{}) interface{} {
			return func(x interface{}) interface{} {
				switch f := f.(type) {
				case func(func(f func(x interface{}) func(x interface{}) interface{}) interface{}) func(f func(x interface{}) func(x interface{}) interface{}) interface{}:
					return f(n(f)(x.(func(f func(x interface{}) func(x interface{}) interface{}) interface{})).(func(f func(x interface{}) func(x interface{}) interface{}) interface{}))
				default:
					return f.(func(x interface{}) interface{})(n(f)(x))
				}
			}
		}
	}(func(n func(f interface{}) func(x interface{}) interface{}) func(f interface{}) func(x interface{}) interface{} {
		return func(f interface{}) func(x interface{}) interface{} {
			return func(x interface{}) interface{} {
				switch f := f.(type) {
				case func(func(f func(x interface{}) func(x interface{}) interface{}) interface{}) func(f func(x interface{}) func(x interface{}) interface{}) interface{}:
					return f(n(f)(x.(func(f func(x interface{}) func(x interface{}) interface{}) interface{})).(func(f func(x interface{}) func(x interface{}) interface{}) interface{}))
				default:
					return f.(func(x interface{}) interface{})(n(f)(x))
				}
			}
		}
	}(func(n func(f interface{}) func(x interface{}) interface{}) func(f interface{}) func(x interface{}) interface{} {
		return func(f interface{}) func(x interface{}) interface{} {
			return func(x interface{}) interface{} {
				switch f := f.(type) {
				case func(func(f func(x interface{}) func(x interface{}) interface{}) interface{}) func(f func(x interface{}) func(x interface{}) interface{}) interface{}:
					return f(n(f)(x.(func(f func(x interface{}) func(x interface{}) interface{}) interface{})).(func(f func(x interface{}) func(x interface{}) interface{}) interface{}))
				default:
					return f.(func(x interface{}) interface{})(n(f)(x))
				}
			}
		}
	}(func(f interface{}) func(x interface{}) interface{} {
		return func(x interface{}) interface{} {
			return x
		}
	}))))).(func(f interface{}) func(x interface{}) interface{})(incrementByOne)(startValue)

	if res != 24 {
		t.Errorf("The factorial of 4 should be 24 instead is %v", res)
	}
}
