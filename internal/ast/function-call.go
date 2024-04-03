package ast

import (
	"github.com/renatopp/langtools/asts"
	"github.com/renatopp/langtools/tokens"
)

type FunctionCall struct {
	Token     tokens.Token
	Target    Identifier
	Arguments []asts.Node
}

func (c FunctionCall) GetToken() tokens.Token {
	return c.Token
}

func (c FunctionCall) String() string {
	return "<function-call>"
}

func (c FunctionCall) Children() []asts.Node {
	children := []asts.Node{c.Target}
	children = append(children, c.Arguments...)
	return children
}
