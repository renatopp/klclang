package main

import (
	"fmt"
	"klc/lang/lexer"
)

var program = `
binsearch(...list, v) {
  m = mid(v)
  ? [empty list] = 0
  ? v > list(m) = [binsearch list(m+1:) v]
  ? v < list(m) = [binsearch list(:m) v]
  ? 1
}
`

func main() {
	lex := lexer.New(program)

	tokens := lex.All()
	for _, token := range tokens {
		fmt.Println(token.ToString())
	}
}
