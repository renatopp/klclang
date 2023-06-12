package main

import (
	"fmt"
	"klc/lang/ast"
	"klc/lang/eval"
	"klc/lang/lexer"
	"klc/lang/parser"
	"strings"
)

var program = `
echo 'Hello, World'
`

func main() {
	input := program

	lex0 := lexer.New(input)
	tokens := lex0.All()
	for _, token := range tokens {
		fmt.Println(token.ToString())
	}

	fmt.Println("--------------")

	lex := lexer.New(input)
	par := parser.New(lex)

	prg, err := par.Parse()
	if err != nil {
		fmt.Println(err)
		return
	}

	prg.Root.Traverse(0, func(level int, node ast.Node) {
		ident := strings.Repeat("  ", level)
		fmt.Println(ident + node.String())
	})

	fmt.Println("--------------")

	eval := eval.New()
	fmt.Println(eval.Eval(prg.Root).AsString())

}
