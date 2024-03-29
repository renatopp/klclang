package ast

import (
	"fmt"

	"github.com/renatopp/langtools"
)

type UnaryOperator struct {
	Token      langtools.Token
	Operator   string
	Expression INode
}

func (u *UnaryOperator) GetToken() langtools.Token {
	return u.Token
}

func (u *UnaryOperator) String() string {
	return fmt.Sprintf("<unary-operator:%s>", u.Operator)
}

func (u *UnaryOperator) Children() []INode {
	return []INode{u.Expression}
}

func (u *UnaryOperator) Traverse(level int, f func(int, INode)) {
	f(level, u)
	u.Expression.Traverse(level+1, f)
}
