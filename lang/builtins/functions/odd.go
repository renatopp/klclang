package functions

import (
	"klc/lang/builtins"
	"klc/lang/obj"
	"math"
)

func odd(ev obj.Evaluator, args ...obj.Object) obj.Object {
	if math.Mod(args[0].AsNumber(), 2) != 0 {
		return builtins.TRUE
	}

	return builtins.FALSE
}

var Odd = builtins.WithDoc(
	builtins.NewFunction(odd, builtins.NewParam("value", nil, false)),
	`Checks if the given 'value' is an odd number.`,
)
