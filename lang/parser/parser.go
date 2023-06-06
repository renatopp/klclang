package parser

import (
	"fmt"
	"klc/lang/ast"
	"klc/lang/lexer"
	"klc/lang/parser/order"
	"klc/lang/token"
	"strconv"
)

var priorities = map[string]int{
	"+":  order.Addition,
	"-":  order.Subtraction,
	"*":  order.Multiplication,
	"/":  order.Division,
	"%":  order.Mod,
	"**": order.Exponentiation,
	"//": order.Division,
	"==": order.Comparison,
	"!=": order.Comparison,
	">":  order.Comparison,
	"<":  order.Comparison,
	">=": order.Comparison,
	"<=": order.Comparison,
	"!":  order.Not,
	"&&": order.And,
	"||": order.Or,
	"++": order.Concat,
}

func priorityOf(t *token.Token) int {
	switch t.Type {
	case token.Operator:
		op := t.Literal
		v, ok := priorities[op]
		if !ok {
			return order.Lowest
		}
		return v
	case token.Assignment:
		return order.Assign
	case token.Arrow:
		return order.Arrow
	case token.Spread:
		return order.Spread
	default:
		return order.Lowest
	}
}

type prefixFn func() ast.Node
type infixFn func(ast.Node) ast.Node

type Parser struct {
	lexer    *lexer.Lexer
	root     *ast.Program
	prefixes map[token.Type]prefixFn
	infixes  map[token.Type]infixFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		lexer:    l,
		root:     &ast.Program{},
		prefixes: make(map[token.Type]prefixFn),
		infixes:  make(map[token.Type]infixFn),
	}

	p.prefixes[token.Number] = p.parsePrefixNumber
	p.infixes[token.Operator] = p.parseInfixOperator

	return p
}

func (p *Parser) Parse() ast.Node {
	return p.parseExpression(order.Lowest)
}

func (p *Parser) expect(t token.Type) {
	if !p.lexer.Peek().Is(t) {
		panic("expected " + string(t) + ", got " + p.lexer.Peek().ToString())
	}
}

func (p *Parser) parseExpression(priority int) ast.Node {
	t := p.lexer.Current()

	for t.Is(token.Newline) {
		fmt.Println("prefix ignoring", t.ToString())
		t = p.lexer.Next()
	}

	fmt.Println("prefix", t.ToString())
	prefix := p.prefixes[t.Type]
	if prefix == nil {
		panic("no prefix function for " + t.ToString())
	}

	root := prefix()

	nt := p.lexer.Current()
	for nt.Is(token.Newline) {
		fmt.Println("infix ignoring", t.ToString())
		nt = p.lexer.Next()
	}
	for !isEndOfExpression(nt) && priorityOf(nt) >= priority {
		infix := p.infixes[nt.Type]
		if infix == nil {
			return root
		}

		fmt.Println("infixing", nt.ToString())
		root = infix(root)
		nt = p.lexer.Current()
	}

	return root
}

func (p *Parser) parsePrefixNumber() ast.Node {
	t := p.lexer.Current()
	v, e := strconv.ParseFloat(t.Literal, 64)

	if e != nil {
		panic(e)
	}

	p.lexer.Next()
	return &ast.Number{
		Token: t,
		Value: v,
	}
}

func (p *Parser) parseInfixOperator(left ast.Node) ast.Node {
	t := p.lexer.Current()
	pr := priorityOf(t)

	node := &ast.BinaryOperation{
		Token: t,
	}

	p.lexer.Next()

	node.Left = left
	node.Right = p.parseExpression(pr)

	return node
}

func isEndOfExpression(t *token.Token) bool {
	return t.Is(token.Semicolon) // t.Is(token.Newline) ||
}
