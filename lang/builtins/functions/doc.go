package functions

import (
	"klc/lang/builtins"
	"klc/lang/obj"
)

func doc(args ...obj.Object) obj.Object {
	d := args[0].GetDoc()
	if d == "" {
		d = "No documentation available."
	}
	return builtins.NewString(d)
}

var Doc = builtins.WithDoc(
	builtins.NewFunction(doc, builtins.NewParam("var", nil, false)),
	`Returns the documentation for a given variable.`,
)
