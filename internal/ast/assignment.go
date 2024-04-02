package ast

import (
	"fmt"

	"github.com/renatopp/langtools/asts"
	"github.com/renatopp/langtools/tokens"
)

type Assignment struct {
	Token      tokens.Token
	Operator   string
	Identifier asts.Node
	Expression asts.Node
}

func (a Assignment) GetToken() tokens.Token {
	return a.Token
}

func (a Assignment) String() string {
	return fmt.Sprintf("<assignment:%s>", a.Operator)
}

func (a Assignment) Children() []asts.Node {
	return []asts.Node{a.Identifier, a.Expression}
}
