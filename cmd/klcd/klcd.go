package main

import (
	"os"

	"github.com/renatopp/klclang/internal"
	"github.com/renatopp/langtools/asts"
)

func main() {
	Debug([]byte(os.Args[1]))
}

func Debug(code []byte) {
	{
		println("CMD:")
		for _, arg := range os.Args {
			println("-", arg)
		}
		println()

		println("CODE:", string(code))
		println()
	}

	{
		println("TOKENS:")
		lexer := internal.NewLexer(code)
		tokens := lexer.All()

		for i, t := range tokens {
			println("-", i, t.DebugString())
		}

		println()

		if lexer.HasErrors() {
			println("ERRORS:")
			println(internal.ConvertLexerErrors(code, lexer.Errors()).Error())
			println()
			os.Exit(0)
		}
	}

	{
		println("NODES:")
		lexer := internal.NewLexer(code)
		parser := internal.NewParser(lexer)

		node := parser.Parse()
		if parser.HasErrors() {
			println("ERRORS:")
			println(internal.ConvertParserErrors(code, parser.Errors()).Error())
			println()
			os.Exit(0)
		}

		asts.Print(node, "  ")
		println()
	}

	{
		println("EVALUATOR:")
		obj, err := internal.Run(code)
		if err != nil {
			println("ERRORS:")
			println(err.Error())
			println()
			os.Exit(0)
		}

		println("=", obj.String())
		println()
	}
}
