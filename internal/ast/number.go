package ast

import (
	"fmt"

	"github.com/renatopp/langtools/asts"
	"github.com/renatopp/langtools/tokens"
)

type Number struct {
	Token tokens.Token
	Value float64
}

func (n Number) GetToken() tokens.Token {
	return n.Token
}

func (n Number) String() string {
	return fmt.Sprintf("<number:%0.4f>", n.Value)
}

func (n Number) Children() []asts.Node {
	return []asts.Node{}
}
