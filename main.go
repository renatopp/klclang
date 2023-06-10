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
binsearch = fn (...list, v) {
  m = mid(v)
  ? empty list: 0
  ? v > list(m): binsearch list[m+1:], v
  ? v < list(m): binsearch list[:m], v
  ? 1
}
`

var expression = `
quicksort = fn(...list) {
  ? empty(list) : list
  p = mid(list)
  concat(
    (filter list, where x < p),
    (filter list, where x == p),
    (filter list, where x > p)
	)
}
`

var testing2 = `
t = 'hello'

'teste'
	.echo()
	.exit()
`

var testing = `
solution = fn(n=1000) {
  range(n)
    .filter where x%3 == 0 || x%5 == 0
    .sum()
}

assert solution(3) == 0
assert solution(5) == 1
assert solution(6) == 2
`

func main() {
	input := testing2

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
