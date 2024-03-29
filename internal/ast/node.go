package ast

import "github.com/renatopp/langtools"

type INode interface {
	GetToken() langtools.Token
	String() string
	Children() []INode
	Traverse(int, func(int, INode))
}
