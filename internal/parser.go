package internal

import (
	"strconv"

	"github.com/renatopp/klclang/internal/ast"
	"github.com/renatopp/langtools/asts"
	"github.com/renatopp/langtools/parsers"
	"github.com/renatopp/langtools/tokens"
)

type KlcParser struct {
	*parsers.PrattParser
}

func NewParser(lexer *KlcLexer) *KlcParser {
	k := &KlcParser{}
	k.PrattParser = parsers.NewPrattParser(lexer)
	k.PrattParser.IsEndOfExpr = k.isEndOfExpr
	k.PrattParser.PrecedenceFn = k.precedence

	k.RegisterPrefixFn(TNumber, k.prefixNumber)
	k.RegisterPrefixFn(TIdentifier, k.prefixIdentifier)
	k.RegisterPrefixFn(TLParen, k.prefixParen)
	k.RegisterPrefixFn(TOperator, k.prefixOperator)
	k.RegisterInfixFn(TAssign, k.infixAssignment)
	k.RegisterInfixFn(TOperator, k.infixOperator)
	k.RegisterInfixFn(TKeyword, k.infixKeyword)

	return k
}

func (k *KlcParser) isEndOfExpr(t tokens.Token) bool {
	return t.Type == TEoe || t.Type == TEof
}

const PRECEDENCE_UNARY = 36

func (k *KlcParser) precedence(t tokens.Token) int {
	switch {
	case t.IsType(TComma):
		return 10

	case t.IsType(TAssign):
		return 15

	case t.IsType(TOperator) && t.IsOneOfLiterals("or"):
		return 20
	case t.IsType(TOperator) && t.IsOneOfLiterals("and"):
		return 21
	case t.IsType(TOperator) && t.IsOneOfLiterals("!"):
		return 22

	case t.IsType(TOperator) && t.IsOneOfLiterals("==", "!="):
		return 25
	case t.IsType(TOperator) && t.IsOneOfLiterals(">", "<", ">=", "<="):
		return 26

	case t.IsType(TKeyword) && t.IsOneOfLiterals("to"):
		return 29
	case t.IsType(TOperator) && t.IsOneOfLiterals("+", "-"):
		return 30
	case t.IsType(TOperator) && t.IsOneOfLiterals("*", "/"):
		return 31
	case t.IsType(TOperator) && t.IsOneOfLiterals("%"):
		return 32
	case t.IsType(TOperator) && t.IsOneOfLiterals("^"):
		return 34
	// UNARY

	case t.IsType(TLParen):
		return 40

	default:
		return 0
	}
}

func (k *KlcParser) Parse() asts.Node {
	first := k.Lexer.PeekToken()
	statements := []asts.Node{}

	cur := k.Lexer.PeekToken()
	for !cur.IsType(TEof) {
		switch {
		case cur.IsType(TEoe):
			k.Lexer.EatToken()

		default:
			node := k.ParseExpression(0)
			if node == nil {
				k.RegisterError("invalid expression")
				return nil
			}

			statements = append(statements, node)
		}
		cur = k.Lexer.PeekToken()
	}

	return ast.Block{
		Token:      first,
		Statements: statements,
	}
}

func (k *KlcParser) prefixNumber() asts.Node {
	t := k.Lexer.EatToken()
	v, err := strconv.ParseFloat(t.Literal, 64)
	if err != nil {
		k.RegisterErrorWithToken(err.Error(), t)
	}

	return ast.Number{
		Token: t,
		Value: v,
	}
}

func (k *KlcParser) prefixIdentifier() asts.Node {
	t := k.Lexer.EatToken()
	return ast.Identifier{
		Token: t,
		Name:  t.Literal,
	}
}

func (k *KlcParser) prefixParen() asts.Node {
	t := k.Lexer.EatToken()
	expr := k.ParseExpression(0)

	if expr == nil {
		k.RegisterErrorWithToken("expected expression", t)
		return nil
	}

	if !k.ExpectType(TRParen) {
		return nil
	}

	k.Lexer.EatToken()
	return expr
}

func (k *KlcParser) prefixOperator() asts.Node {
	cur := k.Lexer.EatToken()

	if !cur.IsOneOfLiterals("-", "+") {
		k.RegisterErrorWithToken("expected unary operator", cur)
		return nil
	}

	return ast.UnaryOperator{
		Token:      cur,
		Operator:   cur.Literal,
		Expression: k.ParseExpression(PRECEDENCE_UNARY),
	}
}

func (k *KlcParser) infixAssignment(left asts.Node) asts.Node {
	cur := k.Lexer.EatToken()
	right := k.ParseExpression(k.PrecedenceFn(cur))

	if isNodeAn[ast.Identifier](left) {
		return ast.Assignment{
			Token:      cur,
			Operator:   cur.Literal,
			Identifier: left.(ast.Identifier),
			Expression: right,
		}
	}

	k.RegisterErrorWithToken("expected identifier or function definition at left", left.GetToken())
	return nil
}

func (k *KlcParser) infixOperator(left asts.Node) asts.Node {
	cur := k.Lexer.EatToken()
	right := k.ParseExpression(k.PrecedenceFn(cur))

	if right == nil {
		k.RegisterErrorWithToken("expected expression", cur)
		return nil
	}

	return ast.BinaryOperator{
		Token:    cur,
		Operator: cur.Literal,
		Left:     left,
		Right:    right,
	}
}

func (k *KlcParser) infixKeyword(left asts.Node) asts.Node {
	t := k.Lexer.EatToken()
	right := k.ParseExpression(k.PrecedenceFn(t))

	if right == nil {
		k.RegisterErrorWithToken("expected expression", t)
		return nil
	}

	return ast.BinaryOperator{
		Token:    t,
		Operator: "/",
		Left:     left,
		Right:    right,
	}
}

func isNodeAn[T asts.Node](node asts.Node) bool {
	_, ok := node.(T)
	return ok
}
