package functions

import (
	"klc/lang/builtins"
	"klc/lang/obj"
)

func boolean(args ...obj.Object) obj.Object {
	if args[0].AsBool() {
		return builtins.TRUE
	}

	return builtins.FALSE
}

var Boolean = builtins.WithDoc(
	builtins.NewFunction(boolean, builtins.NewParam("value", nil, false)),
	`Converts the value to a boolean (0 or 1).`,
)
