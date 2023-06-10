package functions

// import (
// 	"klc/lang/builtins"
// 	"klc/lang/obj"
// )

// func filter(ev obj.Evaluator, args ...obj.Object) obj.Object {
// 	if len(args) != 2 {
// 		panic("wrong number of arguments")
// 	}

// 	sequence := args[0].(*obj.List)
// 	fn, ok := args[1].(obj.Callable)
// 	if !ok {
// 		panic("second argument must be a function")
// 	}

// 	results := make([]obj.Object, 0)
// 	args = make([]obj.Object, 1)
// 	for _, x := range sequence.Values {
// 		args[0] = x
// 		r := ev.Call(fn.(obj.Callable), args)
// 		if r.AsBool() {
// 			results = append(results, x)
// 		}
// 	}

// 	return builtins.NewList(results...)
// }

// var Filter = builtins.WithDoc(
// 	builtins.NewFunction(filter,
// 		builtins.NewParam("list", nil, true),
// 		builtins.NewParam("fn", nil, false),
// 	),
// 	`Returns a new list containing only the elements of the list for which the function returns true.`,
// )
