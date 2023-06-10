package functions

import (
	"fmt"
	"klc/lang/builtins"
	"klc/lang/obj"
)

func assert(ev obj.Evaluator, args ...obj.Object) obj.Object {
	msg := "invalid assertion"
	if len(args) > 1 {
		msg = args[1].AsString()
	}

	if !args[0].AsBool() {
		fmt.Println("ERR!", msg)
		exit(ev, builtins.NewNumber(1))
		return builtins.FALSE
	}

	return builtins.TRUE
}

var Assert = builtins.WithDoc(
	builtins.NewFunction(assert,
		builtins.NewParam("value", nil, false),
		builtins.NewParam("message", builtins.NewString("invalid assertion"), false),
	),
	`Test if the given value is true. Exiting the program if it is not.`,
)
