package main

import (
	"os"
	"strings"

	"github.com/renatopp/klclang/cmd/klc/cmds"
)

func main() {
	if len(os.Args) > 1 {
		cmds.Run([]byte(strings.Join(os.Args[1:], " ")))
	} else {
		cmds.Repl()
	}
}
