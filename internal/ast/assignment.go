package ast

import (
	"fmt"

	"github.com/renatopp/langtools/asts"
	"github.com/renatopp/langtools/tokens"
)

type Assignment struct {
	Token         tokens.Token
	Operator      string
	Documentation string
	Identifier    *Identifier
	Expression    asts.Node
}

func (a Assignment) GetToken() tokens.Token {
	return a.Token
}

func (a Assignment) String() string {
	docs := ""
	if a.Documentation != "" {
		docs = fmt.Sprintf(" -- %s", a.Documentation)
	}
	return fmt.Sprintf("<assignment:%s>%s", a.Operator, printRaw(docs))
}

func (a Assignment) Children() []asts.Node {
	return []asts.Node{a.Identifier, a.Expression}
}
