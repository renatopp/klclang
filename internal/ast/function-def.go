package ast

import (
	"fmt"

	"github.com/renatopp/langtools/asts"
	"github.com/renatopp/langtools/tokens"
)

type FunctionDef struct {
	Token         tokens.Token
	Name          string
	Documentation string
	Params        []asts.Node
	Body          asts.Node
}

func (d FunctionDef) GetToken() tokens.Token {
	return d.Token
}

func (d FunctionDef) String() string {
	docs := ""
	if d.Documentation != "" {
		docs = fmt.Sprintf(" -- %s", d.Documentation)
	}
	return fmt.Sprintf("<function-def:%s>%s", d.Name, printRaw(docs))
}

func (d FunctionDef) Children() []asts.Node {
	return append(append([]asts.Node{}, d.Params...), d.Body)
}
