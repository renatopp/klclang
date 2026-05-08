package main

import "github.com/renatopp/go-cli"

func main() {
	cli.Name("klc")
	cli.Description("Command line calculator for your daily use.")
	cli.Version("0.2.0")
	cli.AutoHelp(true)

	lex := cli.FlagBool("lex", "l", "[dev only] Show lexical output")
	parse := cli.FlagBool("parse", "p", "[dev only] Show parsing output")

	args := cli.Pos("exprs", "Expressions to evaluate").AsVariadic()
	cli.Parse()

	if len(args.Values()) == 0 {
		println("This is a REPL mock")

	} else {
		println("Evaluating expressions:")
		for _, expr := range args.Values() {
			println(" -", expr)
		}
	}

	println("Lex:", lex.Value())
	println("Parse:", parse.Value())
}
