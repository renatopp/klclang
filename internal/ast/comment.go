package ast

import (
	"fmt"

	"github.com/renatopp/langtools/asts"
	"github.com/renatopp/langtools/tokens"
)

type Comment struct {
	Token tokens.Token
	Name  string
}

func (n Comment) GetToken() tokens.Token {
	return n.Token
}

func (n Comment) String() string {
	return fmt.Sprintf("<comment:%s>", n.Name)
}

func (n Comment) Children() []asts.Node {
	return []asts.Node{}
}
