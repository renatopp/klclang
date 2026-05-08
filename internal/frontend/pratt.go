package frontend

import (
	"github.com/renatopp/klclang/internal/core"
)

type peekFn func() *core.Token
type prefixFn func() core.Node
type infixFn func(core.Node) core.Node
type postfixFn func(core.Node) core.Node

type PrattSolver struct {
	peek        peekFn
	precedences map[core.TokenKind]int
	prefixFns   map[core.TokenKind]prefixFn
	infixFns    map[core.TokenKind]infixFn
	postfixFns  map[core.TokenKind]postfixFn
}

func NewPrattSolver(peek peekFn) *PrattSolver {
	return &PrattSolver{
		peek:        peek,
		prefixFns:   map[core.TokenKind]prefixFn{},
		infixFns:    map[core.TokenKind]infixFn{},
		postfixFns:  map[core.TokenKind]postfixFn{},
		precedences: map[core.TokenKind]int{},
	}
}

func (p *PrattSolver) RegisterPrefixFn(kind core.TokenKind, fn prefixFn) *PrattSolver {
	p.prefixFns[kind] = fn
	return p
}

func (p *PrattSolver) RegisterInfixFn(kind core.TokenKind, fn infixFn) *PrattSolver {
	p.infixFns[kind] = fn
	return p
}

func (p *PrattSolver) RegisterPostfixFn(kind core.TokenKind, fn postfixFn) *PrattSolver {
	p.postfixFns[kind] = fn
	return p
}

func (p *PrattSolver) RegisterPrecedence(kind core.TokenKind, val int) *PrattSolver {
	p.precedences[kind] = val
	return p
}

func (p *PrattSolver) PrecedenceOf(kind core.TokenKind) int {
	if val, ok := p.precedences[kind]; ok {
		return val
	}
	return 0
}

func (p *PrattSolver) Prefix() core.Node {
	prefix := p.prefixFns[p.peek().Kind]
	if prefix == nil {
		return nil
	}
	return prefix()
}

func (p *PrattSolver) Solve(precedence int) core.Node {
	prefix := p.prefixFns[p.peek().Kind]
	if prefix == nil {
		return nil
	}
	left := prefix()
	if left == nil {
		return nil
	}

	cur := p.peek()
	for {
		starting := cur

		if precedence < p.PrecedenceOf(cur.Kind) {
			infix := p.infixFns[cur.Kind]
			if infix != nil {
				left = infix(left)
				cur = p.peek()
			}
		}

		for {
			postfix := p.postfixFns[cur.Kind]
			if postfix == nil {
				break
			}
			newLeft := postfix(left)
			if newLeft == nil {
				break
			}
			left = newLeft
			cur = p.peek()
		}

		// Didn't find any infix or postfix function
		if starting == cur {
			break
		}
	}

	return left
}
