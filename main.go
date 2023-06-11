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
a = fn x, y {
	fn z {
		x + y + z
	}
}

b = a(1, 2)
echo b(3)
`

var testing = `

solution = fn max, a=1, b=1, sum=0 {
	c = a + b
	? c > max: sum
	? solution max, b, c, c is even ? sum + c : sum
}

assert solution(10) == 10
assert solution(50) == 44

solution 4000000
`

func main() {
	input := testing

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
