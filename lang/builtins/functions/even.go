package functions

import (
	"klc/lang/builtins"
	"klc/lang/obj"
	"math"
)

func even(args ...obj.Object) obj.Object {
	if math.Mod(args[0].AsNumber(), 2) == 0 {
		return builtins.TRUE
	}

	return builtins.FALSE
}

var Even = builtins.WithDoc(
	builtins.NewFunction(even, builtins.NewParam("value", nil, false)),
	`Checks if the given 'value' is an even number.`,
)
