package ast

import (
	"github.com/renatopp/langtools/asts"
	"github.com/renatopp/langtools/tokens"
)

type Block struct {
	Token      tokens.Token
	Statements []asts.Node
}

func (d Block) GetToken() tokens.Token {
	return d.Token
}

func (b Block) String() string {
	return "<block>"
}

func (b Block) Children() []asts.Node {
	return b.Statements
}
