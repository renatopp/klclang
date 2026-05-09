package main

import (
	"github.com/renatopp/go-cli"
	"github.com/renatopp/klclang/internal"
	"github.com/renatopp/klclang/internal/core"
	"github.com/renatopp/x/fsx"
)

func main() {
	cli.Name("klc")
	cli.Description("Command line calculator for your daily use.")
	cli.Version("0.2.0")
	cli.AutoHelp(true)

	lex := cli.FlagBool("lex", "l", "[dev only] Show lexical output")
	parse := cli.FlagBool("parse", "p", "[dev only] Show parsing output")

	args := cli.Pos("exprs", "Expressions to evaluate").AsVariadic()
	cli.Parse()

	runner := internal.NewRunner()
	runner.PrintLex = lex.Value()
	runner.PrintParse = parse.Value()

	if len(args.Values()) == 0 {
		println("This is a REPL mock")

	} else {
		evaluate(runner, args.Values())
	}
}

func evaluate(runner *internal.Runner, exprs []string) {
	var result any
	var err error
	for _, expr := range exprs {
		if fsx.IsFile(expr) {
			result, err = runner.Run(expr)
		} else {
			result, err = runner.Eval([]byte(expr))
		}

		if err != nil {
			core.PPrintError(err)
			continue
		} else {
			println("Result:", result)
		}
	}
}
