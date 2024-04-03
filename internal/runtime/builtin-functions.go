package runtime

import "math"

var (
	FN_FLOOR = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		// check args
		n := args[0].Number()
		return NewNumber(math.Floor(n))
	}), "Floor rounds a number down to the nearest integer.")

	FN_ROUND = withDocs(NewBuiltinFunction(func(env *Scope, args ...Object) Object {
		// check args
		n := args[0].Number()
		return NewNumber(math.Round(n))
	}), "Round rounds a number to the nearest integer.")
)

func registerFunctions(scope *Scope) {
	scope.Set("floor", FN_FLOOR)
	scope.Set("round", FN_ROUND)
}
