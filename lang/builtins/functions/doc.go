package functions

import (
	"klc/lang/builtins"
	"klc/lang/obj"
)

func doc(args ...obj.Object) obj.Object {
	if len(args) > 1 && args[1].AsNumber() != -1 {
		d := args[1].AsString()
		args[0].SetDoc(d)
		return builtins.NewString(d)
	}

	d := args[0].GetDoc()
	if d == "" {
		d = "No documentation available."
	}
	return builtins.NewString(d)
}

var Doc = builtins.WithDoc(
	builtins.NewFunction(doc,
		builtins.NewParam("var", nil, false),
		builtins.NewParam("doc", builtins.NewNumber(-1), false),
	),
	`Returns the documentation for a given variable, or sets a new documentation for the variable.`,
)
