package functions

import (
	"klc/lang/builtins"
	"klc/lang/obj"
)

func len_(ev obj.Evaluator, args ...obj.Object) obj.Object {
	if len(args) < 1 {
		panic("wrong number of arguments")
	}

	switch v := args[0].(type) {
	case *obj.List:
		return builtins.NewNumber(float64(len(v.Values)))
	case *obj.String:
		return builtins.NewNumber(float64(len(v.Value)))
	default:
		return builtins.NewNumber(1)
	}
}

var Len = builtins.WithDoc(
	builtins.NewFunction(len_,
		builtins.NewParam("list", nil, false),
	),
	`.`,
)
