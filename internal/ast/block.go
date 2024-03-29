package ast

import "github.com/renatopp/langtools"

type Block struct {
	Statements []INode
}

func (b *Block) GetToken() langtools.Token {
	return langtools.Token{}
}

func (b *Block) String() string {
	return "<block>"
}

func (b *Block) Children() []INode {
	return b.Statements
}

func (b *Block) Traverse(level int, f func(int, INode)) {
	f(level, b)
	for _, s := range b.Statements {
		s.Traverse(level+1, f)
	}
}
