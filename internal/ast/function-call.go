package ast

import "github.com/renatopp/langtools"

type FunctionCall struct {
	Target    INode
	Arguments []INode
}

func (c *FunctionCall) GetToken() langtools.Token {
	return c.Target.GetToken()
}

func (c *FunctionCall) String() string {
	return "<function-call>"
}

func (c *FunctionCall) Children() []INode {
	children := []INode{c.Target}
	children = append(children, c.Arguments...)
	return children
}

func (c *FunctionCall) Traverse(level int, f func(int, INode)) {
	f(level, c)
	f(level+1, c.Target)
	for _, a := range c.Arguments {
		a.Traverse(level+1, f)
	}
}
