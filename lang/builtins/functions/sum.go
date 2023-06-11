package functions

import (
	"klc/lang/builtins"
	"klc/lang/obj"
)

func sum(ev obj.Evaluator, args ...obj.Object) obj.Object {
	if len(args) < 1 {
		panic("wrong number of arguments")
	}

	sequence := args[0].(*obj.List)

	var s float64
	for _, v := range sequence.Values {
		s += v.AsNumber()
	}

	return builtins.NewNumber(s)
}

var Sum = builtins.WithDoc(
	builtins.NewFunction(sum,
		builtins.NewParam("list", nil, true),
	),
	`Sum elements of the list.`,
)
