package main

import (
	"os"

	"github.com/renatopp/klclang/internal"
)

func main() {
	if len(os.Args) != 2 {
		println("Usage: klcc <file>")
		os.Exit(1)
	}

	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	obj, err := internal.Run(file)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	println(obj.String())
}
