package ast

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

type AssignmentStatement struct {
	Name *Identifier
	Value Expression
}

func (as *AssignmentStatement) statementNode() {}

func (as *AssignmentStatement) TokenLiteral() string {
	return as.Name.TokenLiteral()
}

type Identifier struct {
	Token tokenizer.Token
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}