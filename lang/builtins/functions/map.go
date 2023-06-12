package functions

import (
	"klc/lang/builtins"
	"klc/lang/obj"
)

func map_(ev obj.Evaluator, args ...obj.Object) obj.Object {
	if len(args) != 2 {
		panic("wrong number of arguments")
	}

	sequence := args[0].(*obj.List)
	fn, ok := args[1].(obj.Callable)
	if !ok {
		panic("second argument must be a function")
	}

	results := make([]obj.Object, 0)
	args = make([]obj.Object, 2)
	for i, x := range sequence.Values {
		args[0] = x
		args[1] = builtins.NewNumber(float64(i))
		r := ev.Call(fn.(obj.Callable), args)
		results = append(results, r)
	}

	return builtins.NewList(results...)
}

var Map = builtins.WithDoc(
	builtins.NewFunction(map_,
		builtins.NewParam("list", nil, true),
		builtins.NewParam("fn", nil, false),
	),
	`Returns a new list containing the results of applying the function to each element of the list.`,
)
