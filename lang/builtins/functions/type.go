package functions

import (
	"klc/lang/builtins"
	"klc/lang/obj"
)

func type_(args ...obj.Object) obj.Object {
	return builtins.NewString(string(args[0].Type()))
}

var Type = builtins.WithDoc(
	builtins.NewFunction(type_, builtins.NewParam("value", nil, false)),
	`Returns the type of the value.`,
)
