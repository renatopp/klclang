package ast

import (
	"fmt"

	"github.com/renatopp/langtools"
)

type Number struct {
	Token langtools.Token
	Value float64
}

func (n *Number) GetToken() langtools.Token {
	return n.Token
}

func (n *Number) String() string {
	return fmt.Sprintf("<number:%0.4f>", n.Value)
}

func (n *Number) Children() []INode {
	return []INode{}
}

func (n *Number) Traverse(level int, f func(int, INode)) {
	f(level, n)
}
