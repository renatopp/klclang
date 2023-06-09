package functions

import (
	"fmt"
	"klc/lang/builtins"
	"klc/lang/obj"
	"strings"
)

func echo(args ...obj.Object) obj.Object {
	b := strings.Builder{}

	for i, x := range args {
		if i > 0 {
			b.WriteString(" ")
		}
		b.WriteString(x.AsString())
	}

	v := b.String()
	fmt.Println(v)
	return builtins.FALSE
}

var Echo = builtins.WithDoc(
	builtins.NewFunction(echo, builtins.NewParam("s", nil, true)),
	`Prints provided args on screen.`,
)
