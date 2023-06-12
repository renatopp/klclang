package functions

import (
	"klc/lang/builtins"
	"klc/lang/obj"
)

func reduce_(ev obj.Evaluator, args ...obj.Object) obj.Object {
	if len(args) < 2 {
		panic("wrong number of arguments")
	}

	sequence := args[0].(*obj.List)
	fn, ok := args[1].(obj.Callable)
	if !ok {
		panic("second argument must be a function")
	}
	var initial obj.Object
	initial = builtins.NewNumber(0)
	if len(args) == 2 {
		initial = args[3]
	}

	args = make([]obj.Object, 2)
	acc := initial
	for _, x := range sequence.Values {
		args[0] = acc
		args[1] = x
		acc = ev.Call(fn.(obj.Callable), args)
	}

	return acc
}

var Reduce = builtins.WithDoc(
	builtins.NewFunction(reduce_,
		builtins.NewParam("list", nil, true),
		builtins.NewParam("fn", nil, false),
		builtins.NewParam("acc", builtins.NewNumber(0), false),
	),
	``,
)
