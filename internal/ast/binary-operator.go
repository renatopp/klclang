package ast

import (
	"fmt"

	"github.com/renatopp/langtools/asts"
	"github.com/renatopp/langtools/tokens"
)

type BinaryOperator struct {
	Token    tokens.Token
	Operator string
	Left     asts.Node
	Right    asts.Node
}

func (b BinaryOperator) GetToken() tokens.Token {
	return b.Token
}

func (b BinaryOperator) String() string {
	return fmt.Sprintf("<binary-operator:%s>", b.Operator)
}

func (b BinaryOperator) Children() []asts.Node {
	return []asts.Node{b.Left, b.Right}
}
