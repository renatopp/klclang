package main

import "github.com/renatopp/klclang/internal"

func main() {
	code := []byte(`2`)
	lexer := internal.NewLexer(code)
	parser := internal.NewParser(lexer)
	node := parser.Parse()
	obj := internal.Eval(node)
	println(obj.String())
}
