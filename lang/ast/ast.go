package ast

type Node interface{}
type Statement interface{}
type Expression interface{}

type Program struct {
	Statements []Statement
}

func (p *Program) Statement() {}
