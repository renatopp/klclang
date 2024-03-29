package ast

import (
	"fmt"

	"github.com/renatopp/langtools"
)

type Identifier struct {
	Token langtools.Token
	Value string
}

func (i *Identifier) GetToken() langtools.Token {
	return i.Token
}

func (i *Identifier) String() string {
	return fmt.Sprintf("<identifier:%s>", i.Value)
}

func (i *Identifier) Children() []INode {
	return []INode{}
}

func (i *Identifier) Traverse(level int, f func(int, INode)) {
	f(level, i)
}
