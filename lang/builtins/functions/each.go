package functions

import (
	"klc/lang/builtins"
	"klc/lang/obj"
)

func each_(ev obj.Evaluator, args ...obj.Object) obj.Object {
	if len(args) != 2 {
		panic("wrong number of arguments")
	}

	sequence := args[0].(*obj.List)
	fn, ok := args[1].(obj.Callable)
	if !ok {
		panic("second argument must be a function")
	}

	args = make([]obj.Object, 2)
	for i, x := range sequence.Values {
		args[0] = x
		args[1] = builtins.NewNumber(float64(i))
		ev.Call(fn.(obj.Callable), args)
	}

	return sequence
}

var Each = builtins.WithDoc(
	builtins.NewFunction(each_,
		builtins.NewParam("list", nil, true),
		builtins.NewParam("fn", nil, false),
	),
	`Returns the same list but the function is applied to each element.`,
)
