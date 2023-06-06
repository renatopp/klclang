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
}

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

type BinaryOperation struct {
	Token *token.Token
	Left  Node
	Right Node
}

func (n *BinaryOperation) String() string {
	return fmt.Sprintf("%s", n.Token.ToString())
}

func (n *BinaryOperation) Traverse(level int, fn func(int, Node)) {
	fn(level, n)
	n.Left.Traverse(level+1, fn)
	n.Right.Traverse(level+1, fn)
}

type UnaryOperation struct {
	Token *token.Token
	Right Node
}

func (n *UnaryOperation) String() string {
	return fmt.Sprintf("%s", n.Token.ToString())
}

func (n *UnaryOperation) Traverse(level int, fn func(int, Node)) {
	fn(level, n)
	n.Right.Traverse(level+1, fn)
}
