package main

import (
	"fmt"
	"klc/lang/ast"
	"klc/lang/lexer"
	"klc/lang/parser"
	"strings"
)

var program = `
binsearch = fn (...list, v) {
  m = mid(v)
  ? empty list: 0
  ? v > list(m): binsearch list[m+1:], v
  ? v < list(m): binsearch list[:m], v
  ? 1
}
`

var expression = `

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

	prg := par.Parse()
	prg.Root.Traverse(0, func(level int, node ast.Node) {
		ident := strings.Repeat("  ", level)
		fmt.Println(ident + node.String())
	})

}
