package ast

import (
	"fmt"

	"github.com/renatopp/langtools"
)

type BinaryOperator struct {
	Token    langtools.Token
	Operator string
	Left     INode
	Right    INode
}

func (b *BinaryOperator) GetToken() langtools.Token {
	return b.Token
}

func (b *BinaryOperator) String() string {
	return fmt.Sprintf("<binary-operator:%s>", b.Operator)
}

func (b *BinaryOperator) Children() []INode {
	return []INode{b.Left, b.Right}
}

func (b *BinaryOperator) Traverse(level int, f func(int, INode)) {
	f(level, b)
	b.Left.Traverse(level+1, f)
	b.Right.Traverse(level+1, f)
}
