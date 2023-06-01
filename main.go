package main

import (
	"fmt"
	"klc/pkg/lang"
	"klc/pkg/lang/tokens"
)

func main() {
	token := lang.Token{
		Name:    tokens.Identifier,
		Literal: "foo",
	}

	fmt.Println(token.ToString())
}
