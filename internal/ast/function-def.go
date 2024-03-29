package ast

import (
	"fmt"

	"github.com/renatopp/langtools"
)

type FunctionDef struct {
	Name   string
	Params []INode
	Body   INode
}

func (d *FunctionDef) GetToken() langtools.Token {
	return langtools.Token{}
}

func (d *FunctionDef) String() string {
	return fmt.Sprintf("<function-def:%s>", d.Name)
}

func (d *FunctionDef) Children() []INode {
	return append(append([]INode{}, d.Params...), d.Body)
}

func (d *FunctionDef) Traverse(level int, f func(int, INode)) {
	f(level, d)
	d.Body.Traverse(level+1, f)
	for _, p := range d.Params {
		p.Traverse(level+1, f)
	}
}
