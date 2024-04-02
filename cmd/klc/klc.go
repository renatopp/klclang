package main

import (
	"os"

	"github.com/renatopp/klclang/cmd/klc/cmds"
)

func main() {
	if len(os.Args) > 1 {
		cmds.Debug([]byte(os.Args[1]))
	} else {
		cmds.Repl()
	}
}
