package ast

import (
	"fmt"

	"github.com/renatopp/langtools"
)

type Assignment struct {
	Token      langtools.Token
	Operator   string
	Identifier INode
	Expression INode
}

func (a *Assignment) GetToken() langtools.Token {
	return a.Token
}

func (a *Assignment) String() string {
	return fmt.Sprintf("<assignment:%s", a.Operator)
}

func (a *Assignment) Children() []INode {
	return []INode{a.Identifier, a.Expression}
}

func (a *Assignment) Traverse(level int, f func(int, INode)) {
	f(level, a)
	a.Identifier.Traverse(level+1, f)
	a.Expression.Traverse(level+1, f)
}
