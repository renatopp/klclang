package functions

import (
	"klc/lang/builtins"
	"klc/lang/obj"
	"os"
)

func exit(args ...obj.Object) obj.Object {
	code := 0

	if len(args) > 0 {
		code = int(args[0].AsNumber())
	}

	os.Exit(code)
	return nil
}

var Exit = builtins.WithDoc(
	builtins.NewFunction(exit,
		builtins.NewParam("value", builtins.FALSE, false),
	),
	`Force exit of the process.`,
)
