package ast

import (
	"fmt"
	"klc/lang/token"
)

type Node interface {
	String() string
	Traverse(int, func(int, Node))
}

type Program struct {
	Root Node
}

// ----------------------------------------------------------------------------
// BLOCK
// ----------------------------------------------------------------------------
type Block struct {
	Statements []Node
}

func (n *Block) String() string {
	return fmt.Sprintf("<block>")
}

func (n *Block) Traverse(level int, fn func(int, Node)) {
	fn(level, n)
	for _, s := range n.Statements {
		s.Traverse(level+1, fn)
	}
}

// ----------------------------------------------------------------------------
// IDENTIFIER
// ----------------------------------------------------------------------------
type Identifier struct {
	Token *token.Token
	Value string
}

func (n *Identifier) String() string {
	return fmt.Sprintf("%s (%s)", n.Token.ToString(), n.Value)
}

func (n *Identifier) Traverse(level int, fn func(int, Node)) {
	fn(level, n)
}

// ----------------------------------------------------------------------------
// NUMBER
// ----------------------------------------------------------------------------
type Number struct {
	Token *token.Token
	Value float64
}

func (n *Number) String() string {
	return fmt.Sprintf("%s (%f)", n.Token.ToString(), n.Value)
}

func (n *Number) Traverse(level int, fn func(int, Node)) {
	fn(level, n)
}

// ----------------------------------------------------------------------------
// STRING
// ----------------------------------------------------------------------------
type String struct {
	Token *token.Token
	Value string
}

func (n *String) String() string {
	return fmt.Sprintf("%s (%s)", n.Token.ToString(), n.Value)
}

func (n *String) Traverse(level int, fn func(int, Node)) {
	fn(level, n)
}

// ----------------------------------------------------------------------------
// LIST
// ----------------------------------------------------------------------------
type List struct {
	Values []Node
}

func (n *List) String() string {
	return fmt.Sprintf("<list>")
}

func (n *List) Traverse(level int, fn func(int, Node)) {
	fn(level, n)
	for _, v := range n.Values {
		v.Traverse(level+1, fn)
	}
}

// ----------------------------------------------------------------------------
// BINARYOPERATION
// ----------------------------------------------------------------------------
type BinaryOperation struct {
	Token *token.Token
	Left  Node
	Right Node
}

func (n *BinaryOperation) String() string {
	return fmt.Sprintf("<binary operator> %s", n.Token.ToString())
}

func (n *BinaryOperation) Traverse(level int, fn func(int, Node)) {
	fn(level, n)
	n.Left.Traverse(level+1, fn)
	n.Right.Traverse(level+1, fn)
}

// ----------------------------------------------------------------------------
// UNARYOPERATION
// ----------------------------------------------------------------------------
type UnaryOperation struct {
	Token *token.Token
	Right Node
}

func (n *UnaryOperation) String() string {
	return fmt.Sprintf("<unary operator> %s", n.Token.ToString())
}

func (n *UnaryOperation) Traverse(level int, fn func(int, Node)) {
	fn(level, n)
	n.Right.Traverse(level+1, fn)
}

// ----------------------------------------------------------------------------
// ASSIGNMENT
// ----------------------------------------------------------------------------
type Assignment struct {
	Token      *token.Token
	Identifier Node
	Expression Node
}

func (n *Assignment) String() string {
	return fmt.Sprintf("%s", n.Token.ToString())
}

func (n *Assignment) Traverse(level int, fn func(int, Node)) {
	fn(level, n)
	n.Identifier.Traverse(level+1, fn)
	n.Expression.Traverse(level+1, fn)
}

// ----------------------------------------------------------------------------
// IFRETURN
// ----------------------------------------------------------------------------
type IfReturn struct {
	Condition Node
	Return    Node
}

func (n *IfReturn) String() string {
	return fmt.Sprintf("<if return>")
}

func (n *IfReturn) Traverse(level int, fn func(int, Node)) {
	fn(level, n)
	n.Condition.Traverse(level+1, fn)
	n.Return.Traverse(level+1, fn)
}

// ----------------------------------------------------------------------------
// CONDITIONAL
// ----------------------------------------------------------------------------
type Conditional struct {
	Condition Node
	True      Node
	False     Node
}

func (n *Conditional) String() string {
	return fmt.Sprintf("<if true false>")
}

func (n *Conditional) Traverse(level int, fn func(int, Node)) {
	fn(level, n)
	n.Condition.Traverse(level+1, fn)
	n.True.Traverse(level+1, fn)
	n.False.Traverse(level+1, fn)
}

// ----------------------------------------------------------------------------
// FUNCTIONCALL
// ----------------------------------------------------------------------------
type FunctionCall struct {
	Function  Node
	Arguments []Node
}

func (n *FunctionCall) String() string {
	return fmt.Sprintf("<function call>")
}

func (n *FunctionCall) Traverse(level int, fn func(int, Node)) {
	fn(level, n)
	n.Function.Traverse(level+1, fn)
	for _, arg := range n.Arguments {
		arg.Traverse(level+1, fn)
	}
}

// ----------------------------------------------------------------------------
// INDEX
// ----------------------------------------------------------------------------
type Index struct {
	Target Node
	Value  Node
}

func (n *Index) String() string {
	return fmt.Sprintf("<index>")
}

func (n *Index) Traverse(level int, fn func(int, Node)) {
	fn(level, n)
	n.Target.Traverse(level+1, fn)
	n.Value.Traverse(level+1, fn)
}

// ----------------------------------------------------------------------------
// SLICE
// ----------------------------------------------------------------------------
type Slice struct {
	Target Node
	From   Node
	To     Node
}

func (n *Slice) String() string {
	return fmt.Sprintf("<slice>")
}

func (n *Slice) Traverse(level int, fn func(int, Node)) {
	fn(level, n)
	n.Target.Traverse(level+1, fn)
	n.From.Traverse(level+1, fn)
	n.To.Traverse(level+1, fn)
}

// ----------------------------------------------------------------------------
// FUNCTIONDEF
// ----------------------------------------------------------------------------
type FunctionDef struct {
	Params []Node
	Block  Node
}

func (n *FunctionDef) String() string {
	return fmt.Sprintf("<function def>")
}

func (n *FunctionDef) Traverse(level int, fn func(int, Node)) {
	fn(level, n)
	for _, param := range n.Params {
		param.Traverse(level+1, fn)
	}
	n.Block.Traverse(level+1, fn)
}

// ----------------------------------------------------------------------------
// DEFAULTARG
// ----------------------------------------------------------------------------
type DefaultArg struct {
	Identifier Node
	Value      Node
}

func (n *DefaultArg) String() string {
	return fmt.Sprintf("<arg with default>")
}

func (n *DefaultArg) Traverse(level int, fn func(int, Node)) {
	fn(level, n)
	n.Identifier.Traverse(level+1, fn)
	n.Value.Traverse(level+1, fn)
}

// ----------------------------------------------------------------------------
// SPREADARG
// ----------------------------------------------------------------------------
type SpreadArg struct {
	Identifier Node
}

func (n *SpreadArg) String() string {
	return fmt.Sprintf("<spread arg>")
}

func (n *SpreadArg) Traverse(level int, fn func(int, Node)) {
	fn(level, n)
	n.Identifier.Traverse(level+1, fn)
}

// ----------------------------------------------------------------------------
// CHAIN
// ----------------------------------------------------------------------------
type Chain struct {
	Left  Node
	Right Node
}

func (n *Chain) String() string {
	return fmt.Sprintf("<chain>")
}

func (n *Chain) Traverse(level int, fn func(int, Node)) {
	fn(level, n)
	n.Left.Traverse(level+1, fn)
	n.Right.Traverse(level+1, fn)
}

// ----------------------------------------------------------------------------
// HELPERS
// ----------------------------------------------------------------------------

func TrueCondition() Node {
	return &Number{
		Token: &token.Token{
			Type:    token.Number,
			Literal: "1",
		},
		Value: 1,
	}
}

func Zero() Node {
	return &Number{
		Token: &token.Token{
			Type:    token.Number,
			Literal: "0",
		},
		Value: 0,
	}
}

func Id(v string) Node {
	return &Identifier{
		Token: &token.Token{
			Type:    token.Identifier,
			Literal: v,
		},
		Value: v,
	}
}
