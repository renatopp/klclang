package cmds

import (
	"os"

	"github.com/renatopp/klclang/internal"
)

func Run(code []byte) {
	obj, err := internal.Run(code)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	println(obj.String())
}
