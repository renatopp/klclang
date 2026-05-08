package ast

import (
	"fmt"

	"github.com/renatopp/langtools/asts"
	"github.com/renatopp/langtools/tokens"
)

type UnaryOperator struct {
	Token      tokens.Token
	Operator   string
	Expression asts.Node
}

func (u UnaryOperator) GetToken() tokens.Token {
	return u.Token
}

func (u UnaryOperator) String() string {
	return fmt.Sprintf("<unary-operator:%s>", u.Operator)
}

func (u UnaryOperator) Children() []asts.Node {
	return []asts.Node{u.Expression}
}
