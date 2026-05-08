package ast

import (
	"fmt"

	"github.com/renatopp/langtools/asts"
	"github.com/renatopp/langtools/tokens"
)

type Identifier struct {
	Token tokens.Token
	Name  string
}

func (n Identifier) GetToken() tokens.Token {
	return n.Token
}

func (n Identifier) String() string {
	return fmt.Sprintf("<identifier:%s>", n.Name)
}

func (n Identifier) Children() []asts.Node {
	return []asts.Node{}
}
