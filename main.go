package main

import (
	"fmt"
	"klc/lang/ast"
	"klc/lang/lexer"
	"klc/lang/parser"
	"strings"
)

var program = `
list.sort()
a = x => {}
binsearch(...list, v) {
  m = mid(v)
  ? [empty list] = 0
  ? v > list(m) = [binsearch list(m+1:) v]
  ? v < list(m) = [binsearch list(:m) v]
  ? 1
}
`

// var expression = `1 + 1`

var expression = `1 * add()`

func main() {
	lex0 := lexer.New(expression)
	tokens := lex0.All()
	for _, token := range tokens {
		fmt.Println(token.ToString())
	}

	fmt.Println("--------------")

	lex := lexer.New(expression)
	par := parser.New(lex)

	root := par.Parse()
	root.Traverse(0, func(level int, node ast.Node) {
		ident := strings.Repeat("  ", level)
		fmt.Println(ident + node.String())
	})

}
