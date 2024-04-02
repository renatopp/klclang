package cmds

import (
	"os"

	"github.com/renatopp/klclang/internal"
	"github.com/renatopp/langtools/asts"
)

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
	}

	{
		println("NODES:")
		lexer := internal.NewLexer(code)
		parser := internal.NewParser(lexer)

		node := parser.Parse()
		if node == nil {
			println("- node is nil")
		}

		asts.Print(node, "  ")
	}
}
